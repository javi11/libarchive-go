package libarchive

/*
#cgo CFLAGS: -I${SRCDIR}/src
#cgo darwin LDFLAGS: ${SRCDIR}/libarchive_darwin.a -lstdc++
#cgo windows,amd64 LDFLAGS: ${SRCDIR}/libarchive_windows_amd64.a -lstdc++
#cgo linux,amd64 LDFLAGS: ${SRCDIR}/libarchive_linux_amd64.a -lstdc++
#cgo linux,arm64 LDFLAGS: ${SRCDIR}libarchive_linux_arm64.a -lstdc++
#include <archive.h>
#include <archive_entry.h>
#include <stdlib.h>
*/
import "C"

import (
	"os"
	"path/filepath"
	"syscall"
)

type ArchiveEntry interface {
	Stat() os.FileInfo
	PathName() string
}

type entry struct {
	entry *C.struct_archive_entry
}

type entryInfo struct {
	stat *C.struct_stat
	name string
}

func (h *entry) Stat() os.FileInfo {
	info := &entryInfo{}
	info.stat = C.archive_entry_stat(h.entry)
	info.name = filepath.Base(h.PathName())
	return info
}

func (h *entry) PathName() string {
	name := C.archive_entry_pathname(h.entry)

	return C.GoString(name)
}

func (e *entryInfo) Name() string {
	return e.name
}
func (e *entryInfo) Size() int64 {
	return int64(e.stat.st_size)
}
func (e *entryInfo) Mode() os.FileMode {
	mode := os.FileMode(e.stat.st_mode & 0777)
	switch e.stat.st_mode & syscall.S_IFMT {
	case syscall.S_IFLNK:
		mode |= os.ModeSymlink
	case syscall.S_IFSOCK:
		mode |= os.ModeSocket
	case syscall.S_IFCHR:
		mode |= os.ModeDevice | os.ModeCharDevice
	case syscall.S_IFBLK:
		mode |= os.ModeDevice
	case syscall.S_IFDIR:
		mode |= os.ModeDir
	case syscall.S_IFIFO:
		mode |= os.ModeNamedPipe
	}
	return mode
}
func (e *entryInfo) IsDir() bool {
	return e.stat.st_mode&syscall.S_IFDIR != 0
}
func (e *entryInfo) Sys() interface{} {
	return e.stat
}
