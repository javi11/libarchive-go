package libarchive

/*
#cgo CFLAGS: -I${SRCDIR}/src
#cgo darwin,arm64 LDFLAGS: ${SRCDIR}/libarchive_darwin_arm64.a -lstdc++ -larchive -static
#cgo darwin,amd64 LDFLAGS: ${SRCDIR}/libarchive_darwin_amd64.a -lstdc++ -larchive -static
#cgo windows,amd64 LDFLAGS: ${SRCDIR}/libarchive_windows_amd64.a -lstdc++ -larchive -static
#cgo linux,amd64 LDFLAGS: ${SRCDIR}/libarchive_linux_amd64.a -lstdc++ -larchive -static
#cgo linux,arm64 LDFLAGS: ${SRCDIR}libarchive_linux_arm64.a -lstdc++ -larchive -static
#include <archive.h>
#include <archive_entry.h>
#include <stdlib.h>
#include "libarchive.h"
*/
import "C"
import (
	"io"
	"unsafe"
)

type RarFile struct {
	archive *C.struct_archive
	index   int64 // current reading index
	opened  []*partFile
}

func (r *RarFile) Next() (ArchiveEntry, error) {
	e := &entry{}

	errn := int(C.archive_read_next_header(r.archive, &e.entry))
	err := codeToError(r.archive, errn)
	if err != nil {
		e = nil
	}

	return e, err
}

func (r *RarFile) Read(b []byte) (int, error) {
	n := int(C.archive_read_data(r.archive, unsafe.Pointer(&b[0]), C.size_t(cap(b))))
	if n == 0 {
		return 0, io.EOF
	}
	if n < 0 {
		return 0, codeToError(r.archive, ARCHIVE_FAILED)
	}

	r.index += int64(n)

	return n, nil
}

func (r *RarFile) Seek(offset int64, whence int) (int64, error) {
	n := int64(C.archive_seek_data(r.archive, C.int64_t(offset), C.int(whence)))
	if n < 0 { // err
		err := codeToError(r.archive, ARCHIVE_FAILED)

		return 0, err
	}

	r.index = int64(n)

	return n, nil
}

func (r *RarFile) Size() int {
	return int(C.archive_filter_bytes(r.archive, C.int(0)))
}

func (r *RarFile) Free() error {
	if C.archive_read_free(r.archive) == ARCHIVE_FATAL {
		return ErrArchiveFatal
	}
	return nil
}

func (r *RarFile) Close() error {
	for _, p := range r.opened {
		if err := p.Close(); err != nil {
			return err
		}
	}

	if C.archive_read_close(r.archive) == ARCHIVE_FATAL {
		return ErrArchiveFatal
	}

	return nil
}
