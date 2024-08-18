package libarchive

import (
	"io"
	"io/fs"
	"os"
	"testing"
)

func TestNewFilesOpen(t *testing.T) {
	files := []string{
		"./fixtures/test_read_format_rar_multivolume.part0001.rar.uu",
		"./fixtures/test_read_format_rar_multivolume.part0002.rar.uu",
		"./fixtures/test_read_format_rar_multivolume.part0003.rar.uu",
		"./fixtures/test_read_format_rar_multivolume.part0004.rar.uu",
	}

	fo, err := NewFilesOpen(files, "", &customFs{})
	if err != nil {
		t.Fatalf("Error on creating Archive from a io.Reader:\n %s", err)
	}

	defer func() {
		err := fo.Free()

		if err != nil {
			t.Fatalf("Error on reader Free:\n %s", err)
		}
	}()

	defer func() {
		err := fo.Close()
		if err != nil {
			t.Fatalf("Error on reader Close:\n %s", err)
		}
	}()

	entry, err := fo.Next()
	if err != nil {
		t.Fatalf("got error on reader.Next():\n%s", err)
	}

	name := entry.PathName()

	if name != "ppmd_lzss_conversion_test.txt" {
		t.Fatalf("got wrong entry.PathName():\n%s", name)
	}

	buf := make([]byte, 1024)
	_, err = io.ReadFull(fo, buf)
	if err != nil {
		t.Fatalf("got error on reader.ReadAll():\n%s", err)
	}

	if len(buf) != 1024 {
		t.Fatalf("got wrong len(buf): %d", len(buf))
	}
}

type customFs struct {
}

func (c *customFs) Open(name string) (fs.File, error) {
	return os.Open(name)
}

func (c *customFs) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

func (c *customFs) OpenFile(name string, flag int, perm os.FileMode) (fs.File, error) {
	return os.OpenFile(name, flag, perm)
}
