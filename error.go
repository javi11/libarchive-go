package libarchive

// #include <archive.h>
import "C"
import (
	"errors"
	"fmt"
	"io"
)

const (
	ARCHIVE_EOF    = 1
	ARCHIVE_OK     = 0
	ARCHIVE_RETRY  = -10
	ARCHIVE_WARN   = -20
	ARCHIVE_FAILED = C.ARCHIVE_FAILED
	ARCHIVE_FATAL  = -30
)

var (
	ErrArchiveRetry = errors.New("libarchive: RETRY [operation failed but can be retried]")
	ErrArchiveWarn  = errors.New("libarchive: WARN [success but non-critical error]")
	ErrArchiveFatal = errors.New("libarchive: FATAL [critical error, archive closing]")
)

func codeToError(archive *C.struct_archive, e int) error {
	switch e {
	case ARCHIVE_EOF:
		return io.EOF
	case ARCHIVE_FATAL:
		return fmt.Errorf("libarchive: FATAL [%s]", errorString(archive))
	case ARCHIVE_FAILED:
		return fmt.Errorf("libarchive: FAILED [%s]", errorString(archive))
	case ARCHIVE_RETRY:
		return fmt.Errorf("libarchive: RETRY [%s]", errorString(archive))
	case ARCHIVE_WARN:
		return fmt.Errorf("libarchive: WARN [%s]", errorString(archive))
	}

	return nil
}

func errorString(archive *C.struct_archive) string {
	return C.GoString(C.archive_error_string(archive))
}
