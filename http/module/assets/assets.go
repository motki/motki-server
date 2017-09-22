// Package assets is the web module responsible for serving static assets.
package assets

import "github.com/motki/motki-server/http/route"

type assetsModule struct{}

func New() *assetsModule {
	return &assetsModule{}
}

func (m *assetsModule) Init(mux *route.ServeMux) error {
	fs := newFileServer()
	mux.ServeMux.Handle("/favicon.ico", fs)
	mux.ServeMux.Handle("/browserconfig.xml", fs)
	mux.ServeMux.Handle("/manifest.json", fs)
	mux.ServeMux.Handle("/images/", fs)
	mux.ServeMux.Handle("/styles/", fs)
	mux.ServeMux.Handle("/scripts/", fs)
	mux.ServeMux.Handle("/fonts/", fs)
	return nil
}
