package libarchive

/*
#cgo CFLAGS: -I${SRCDIR}/src
#cgo darwin,arm64 LDFLAGS: ${SRCDIR}/libarchive_darwin_arm64.a -lstdc++ -static
#cgo darwin,amd64 LDFLAGS: ${SRCDIR}/libarchive_darwin_amd64.a -lstdc++ -static
#cgo windows,amd64 LDFLAGS: ${SRCDIR}/libarchive_windows_amd64.a -lstdc++ -static
#cgo linux,amd64 LDFLAGS: ${SRCDIR}/libarchive_linux_amd64.a -lstdc++ -static
#cgo linux,arm64 LDFLAGS: ${SRCDIR}libarchive_linux_arm64.a -lstdc++ -static
#include <archive.h>
#include <archive_entry.h>
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"io"
	"unsafe"
)

//export openCb
func openCb(archive *C.struct_archive, client_data *C.char) C.int {
	fo := (*partFile)(unsafe.Pointer(client_data))

	f, err := fo.fs.Open(fo.filePath)
	if err != nil {
		return ARCHIVE_FAILED
	}

	rSeeker, ok := f.(io.ReadSeekCloser)
	if !ok {
		return ARCHIVE_FAILED
	}

	fo.reader = rSeeker
	fo.buffer = make([]byte, 1024)

	return ARCHIVE_OK
}

//export closeCb
func closeCb(archive *C.struct_archive, client_data *C.char) C.int {
	fo := (*partFile)(unsafe.Pointer(client_data))

	if fo.reader != nil {
		switchCb(archive, client_data, nil)
		fo = nil
	}

	return ARCHIVE_OK
}

//export switchCb
func switchCb(archive *C.struct_archive, client_data *C.char, client_data2 *C.char) C.int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("ERROR: Recovered in switchCb", r)
		}
	}()

	fo1 := (*partFile)(unsafe.Pointer(client_data))
	fo2 := (*partFile)(unsafe.Pointer(client_data2))

	if fo1 != nil && fo1.reader != nil {
		if err := fo1.reader.Close(); err != nil {
			return ARCHIVE_FAILED
		}

		fo1.reader = nil
		fo1.buffer = nil
	}

	if fo2 != nil {
		f, err := fo2.fs.Open(fo2.filePath)
		if err != nil {
			return ARCHIVE_FAILED
		}

		rSeeker, ok := f.(io.ReadSeekCloser)
		if !ok {
			return ARCHIVE_FAILED
		}

		fo2.reader = rSeeker
		fo2.buffer = make([]byte, 1024)

		return ARCHIVE_OK
	}

	return ARCHIVE_OK
}

//export seekCb
func seekCb(archive *C.struct_archive, client_data *C.char, request C.int64_t, whence C.int) C.int64_t {
	reader := (*partFile)(unsafe.Pointer(client_data))
	offset, err := reader.reader.Seek(int64(request), int(whence))
	if err != nil {
		return C.int64_t(0)
	}
	return C.int64_t(offset)
}

//export skipCb
func skipCb(archive *C.struct_archive, client_data *C.char, request C.int64_t) C.int64_t {
	reader := (*partFile)(unsafe.Pointer(client_data))
	oldOffset := reader.index

	offset, err := reader.reader.Seek(int64(request), 1)
	if err != nil {
		return C.int64_t(0)
	}

	reader.index = offset

	return C.int64_t(offset - oldOffset)
}

//export readCb
func readCb(archive *C.struct_archive, client_data *C.char, block unsafe.Pointer) C.size_t {
	reader := (*partFile)(unsafe.Pointer(client_data))
	read, err := reader.reader.Read(reader.buffer)
	if err != nil && err != io.EOF {
		// Set read error
		read = -1
	} else {
		reader.index += int64(read)
	}

	*(*uintptr)(block) = uintptr(unsafe.Pointer(&reader.buffer[0]))

	return C.size_t(read)
}
