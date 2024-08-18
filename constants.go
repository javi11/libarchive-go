package libarchive

/*
#cgo CFLAGS: -I${SRCDIR}/src
#cgo darwin,arm64 LDFLAGS: ${SRCDIR}/bindings_darwin_arm64.a -lstdc++ -static
#cgo darwin,amd64 LDFLAGS: ${SRCDIR}/bindings_darwin_amd64.a -lstdc++ -static
#cgo windows,amd64 LDFLAGS: ${SRCDIR}/bindings_windows_amd64.a -lstdc++ -static
#cgo linux,amd64 LDFLAGS: ${SRCDIR}/bindings_linux_amd64.a -lstdc++ -static
#cgo linux,arm64 LDFLAGS: ${SRCDIR}bindings_linux_arm64.a -lstdc++ -static
#include<archive_entry.h>
*/

import "C"

var (
	FileTypeRegFile = C.AE_IFREG
	FileTypeSymLink = C.AE_IFLNK
	FileTypeSocket  = C.AE_IFSOCK
	FileTypeCharDev = C.AE_IFCHR
	FileTypeBlkDev  = C.AE_IFBLK
	FileTypeDir     = C.AE_IFDIR
	FileTypeFIFO    = C.AE_IFIFO
)
