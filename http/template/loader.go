// +build !release

package template

import "github.com/tyler-sommer/stick"

func newTemplateLoader(config Config) stick.Loader {
	return stick.NewFilesystemLoader(config.ViewsPath)
}
