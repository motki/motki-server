// +build release

package template

import (
	"bytes"
	"io"
	"path/filepath"

	"github.com/tyler-sommer/stick"
)

type bindataTemplate struct {
	name     string
	contents []byte
}

func (t *bindataTemplate) Name() string {
	return t.name
}

func (t *bindataTemplate) Contents() io.Reader {
	return bytes.NewBuffer(t.contents)
}

func newTemplateLoader(config Config) stick.Loader {
	return &assetLoader{viewsPath: config.ViewsPath}
}

type assetLoader struct {
	viewsPath string
}

func (l *assetLoader) Load(name string) (stick.Template, error) {
	res, err := Asset(filepath.Join(l.viewsPath, name))
	if err != nil {
		return nil, err
	}
	return &bindataTemplate{name, res}, nil
}
