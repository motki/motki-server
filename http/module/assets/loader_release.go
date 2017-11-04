// +build release

package assets

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func newFileServer() http.Handler {
	return http.FileServer(bindataFS{"public", http.Dir("public")})
}

type bindataFS struct {
	prefix string

	fallback http.FileSystem
}

func (fs bindataFS) Open(name string) (http.File, error) {
	b, err := Asset(strings.TrimPrefix(name, "/"))
	if err == nil {
		return newFile(name, b), nil
	}
	if strings.Contains(err.Error(), "not found") {
		if fs.fallback != nil {
			return fs.fallback.Open(name)
		}
		return nil, os.ErrNotExist
	}
	return nil, err
}

var defaultTimestamp = time.Now()

type file struct {
	name    string
	dir     bool
	size    int64
	lastMod time.Time

	*bytes.Reader
	io.Closer
}

func newFile(name string, content []byte) *file {
	lastMod := defaultTimestamp
	if info, err := AssetInfo(name); err == nil {
		lastMod = info.ModTime()
	}
	return &file{
		name,
		false,
		int64(len(content)),
		lastMod,

		bytes.NewReader(content),
		ioutil.NopCloser(nil)}
}

func (f *file) Name() string {
	_, name := filepath.Split(f.name)
	return name
}

func (f *file) Mode() os.FileMode {
	mode := os.FileMode(0644)
	if f.dir {
		return mode | os.ModeDir
	}
	return mode
}

func (f *file) ModTime() time.Time {
	return f.lastMod
}

func (f *file) Size() int64 {
	return f.size
}

func (f *file) IsDir() bool {
	return f.Mode().IsDir()
}

func (f *file) Sys() interface{} {
	return nil
}

func (f *file) Readdir(count int) ([]os.FileInfo, error) {
	return nil, errors.New("not a directory")
}

func (f *file) Stat() (os.FileInfo, error) {
	return f, nil
}
