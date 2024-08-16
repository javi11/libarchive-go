package libarchive

/*
#cgo CFLAGS: -I${SRCDIR}/src
#cgo darwin,arm64 LDFLAGS: ${SRCDIR}/libarchive_darwin_arm64.a -lstdc++ -larchive -static
#cgo darwin,amd64 LDFLAGS: ${SRCDIR}/libarchive_darwin_amd64.a -lstdc++ -larchive -static
#cgo windows,amd64 LDFLAGS: ${SRCDIR}/libarchive_windows_amd64.a -lstdc++ -larchive -static
#cgo linux,amd64 LDFLAGS: ${SRCDIR}/libarchive_linux_amd64.a -lstdc++ -larchive -static
#cgo linux,arm64 LDFLAGS: ${SRCDIR}libarchive_linux_arm64.a -lstdc++ -larchive -static
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
