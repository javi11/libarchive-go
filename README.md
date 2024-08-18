# libarchive-go

libarchive go bindings

# Features

- [x] Read from multipart files
- [x] Read from single rar files
- [x] Seek support (limited to what libarchive supports)
- [x] Password protected files support
- [ ] Read from reader

## Installation

```bash
go get github.com/javi11/libarchive-go
```

You will need to have libarchive installed with pkg-config support.

```bash
sudo apt-get install libarchive-dev
```

## Usage

### Read from multipart files

```go
package main

import (
    "fmt"
    "github.com/javi11/libarchive-go"
)

func main() {
  	files := []string{
		"./fixtures/test_read_format_rar_multivolume.part0001.rar.uu",
		"./fixtures/test_read_format_rar_multivolume.part0002.rar.uu",
		"./fixtures/test_read_format_rar_multivolume.part0003.rar.uu",
		"./fixtures/test_read_format_rar_multivolume.part0004.rar.uu",
	}

	fo, err := NewFilesOpen(files)
	if err != nil {
		panic(err)
	}

	defer func() {
		err := fo.Free()

		if err != nil {
			panic(err)
		}
	}()

	defer func() {
		err := fo.Close()
		if err != nil {
			panic(err)
		}
	}()

	entry, err := fo.Next()
	if err != nil {
		panic(err)
	}

	name := entry.PathName()
	fmt.Println(name)

	buf := make([]byte, 1024)
	_, err = io.ReadFull(fo, buf)
	if err != nil {
		panic(err)
	}
}
```

### Read from single rar files

```go
package main

import (
	"fmt"
	"github.com/javi11/libarchive-go"
)

func main() {
	fo, err := NewFileOpen("./fixtures/test_read_format_rar.rar")
	if err != nil {
		panic(err)
	}

	defer func() {
		err := fo.Free()

		if err != nil {
			panic(err)
		}
	}()

	defer func() {
		err := fo.Close()
		if err != nil {
			panic(err)
		}
	}()

	entry, err := fo.Next()
	if err != nil {
		panic(err)
	}

	name := entry.PathName()
	fmt.Println(name)

	buf := make([]byte, 1024)
	_, err = io.ReadFull(fo, buf)
	if err != nil {
		panic(err)
	}
}
```

### Seek support

```go

package main

import (
	"fmt"
	"github.com/javi11/libarchive-go"
)

func main() {
	fo, err := NewFileOpen("./fixtures/test_read_format_rar.rar")
	if err != nil {
		panic(err)
	}

	defer func() {
		err := fo.Free()

		if err != nil {
			panic(err)
		}
	}()

	defer func() {
		err := fo.Close()
		if err != nil {
			panic(err)
		}
	}()

	entry, err := fo.Next()
	if err != nil {
		panic(err)
	}

	name := entry.PathName()
	fmt.Println(name)

	buf := make([]byte, 1024)
	_, err = io.ReadFull(fo, buf)
	if err != nil {
		panic(err)
	}

	_, err = fo.Seek(0, io.SeekStart)
	if err != nil {
		panic(err)
	}

	_, err = io.ReadFull(fo, buf)
	if err != nil {
		panic(err)
	}
}
```

### Custom Fs

```go
package main

import (
	"fmt"
	"github.com/javi11/libarchive-go"
)

func main() {
	fo, err := NewFileOpen("./fixtures/test_read_format_rar.rar")
	if err != nil {
		t.Fatalf("Error on creating Archive from a io.Reader:\n %s", err)
	}

	defer func() {
		err := fo.Free()

		if err != nil {
			panic(err)
		}
	}()

	defer func() {
		err := fo.Close()
		if err != nil {
			panic(err)
		}
	}()

	entry, err := fo.Next()
	if err != nil {
		panic(err)
	}

	name := entry.PathName()
	fmt.Println(name)

	buf := make([]byte, 1024)
	_, err = io.ReadFull(fo, buf)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 1024, len(buf))

	_, err = fo.Seek(0, io.SeekStart)
	if err != nil {
		panic(err)
	}

	_, err = io.ReadFull(fo, buf)
	if err != nil {
		panic(err)
	}
}
```
