// +build !release

package assets

import "net/http"

func newFileServer() http.Handler {
	return http.FileServer(http.Dir("public"))
}
