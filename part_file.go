package libarchive

/*
#cgo pkg-config: libarchive
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
