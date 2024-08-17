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
#include "libarchive.h"
*/
import "C"
import (
	"io/fs"
	"os"
	"unsafe"
)

func NewFilesOpen(
	filePaths []string,
	password string,
	fs fs.FS,
) (*RarFile, error) {
	if fs == nil {
		fs = &defaultFs{}
	}

	arch := C.archive_read_new()

	opened := make([]*partFile, len(filePaths))
	for i, filePath := range filePaths {
		r := partFile{
			archive:      arch,
			fs:           fs,
			filePath:     filePath,
			archiveIndex: i,
		}

		C.read_append_cb_data_binding(r.archive, (*C.char)(unsafe.Pointer(&r)))

		opened[i] = &r
	}

	C.archive_read_support_filter_all(arch)
	C.archive_read_support_format_all(arch)
	C.archive_read_set_seek_callback(arch, (*C.archive_seek_callback)(C.seek_cb_binding))
	C.archive_read_set_open_callback(arch, (*C.archive_open_callback)(C.open_cb_binding))
	C.archive_read_set_close_callback(arch, (*C.archive_close_callback)(C.close_cb_binding))
	C.archive_read_set_read_callback(arch, (*C.archive_read_callback)(C.read_cb_binding))
	C.archive_read_set_skip_callback(arch, (*C.archive_skip_callback)(C.skip_cb_binding))
	C.archive_read_set_switch_callback(arch, (*C.archive_switch_callback)(C.switch_cb_binding))

	if password != "" {
		C.archive_read_add_passphrase(arch, C.CString(password))
	}

	e := C.archive_read_open1(arch)
	err := codeToError(arch, int(e))
	if err != nil {
		C.archive_read_free(arch)

		return nil, err
	}

	return &RarFile{
		archive: arch,
		opened:  opened,
	}, err
}

type defaultFs struct {
}

func (c *defaultFs) Open(name string) (fs.File, error) {
	return os.Open(name)
}

func (c *defaultFs) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

func (c *defaultFs) OpenFile(name string, flag int, perm os.FileMode) (fs.File, error) {
	return os.OpenFile(name, flag, perm)
}
