package libarchive

/*
#cgo CFLAGS: -I${SRCDIR}/src
#cgo darwin,arm64 LDFLAGS: ${SRCDIR}/libarchive_darwin_arm64.a -lstdc++ -larchive -static
#cgo darwin,amd64 LDFLAGS: ${SRCDIR}/libarchive_darwin_amd64.a -lstdc++ -larchive -static
#cgo windows,amd64 LDFLAGS: ${SRCDIR}/libarchive_windows_amd64.a -lstdc++ -larchive -static
#cgo linux,amd64 LDFLAGS: ${SRCDIR}/libarchive_linux_amd64.a -lstdc++ -larchive -static
#cgo linux,arm64 LDFLAGS: ${SRCDIR}libarchive_linux_arm64.a -lstdc++ -larchive -static
#include <archive.h>
#include <stdlib.h>
*/
import "C"
import (
	"io"
	"io/fs"
)

type partFile struct {
	archive      *C.struct_archive
	fs           fs.FS
	filePath     string
	reader       io.ReadSeekCloser
	buffer       []byte
	index        int64
	archiveIndex int
}

func (p *partFile) Close() error {
	if p.reader != nil {
		return p.reader.Close()
	}

	if p.buffer != nil {
		p.buffer = nil
	}

	return nil
}
