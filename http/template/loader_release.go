// +build release

package template

import (
	"bytes"
	"io"

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

func newTemplateLoader(_ Config) stick.Loader {
	return &assetLoader{}
}

type assetLoader struct{}

func (l *assetLoader) Load(name string) (stick.Template, error) {
	res, err := Asset(name)
	if err != nil {
		return nil, err
	}
	return &bindataTemplate{name, res}, nil
}
