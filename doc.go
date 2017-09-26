// Package motki_server is an EVE Online corporation management tool.
//
// Notably, this project contains the motkid source code, which is the
// main motki application server. motkid functions as a remote grpc server
// for clients, as well as a web application.
package motki_server

//go:generate go-bindata -prefix "./public" -pkg template -tags "release" -ignore .DS_Store -o "./http/template/bindata_release.go" ./views/...
//go:generate go-bindata -prefix "./views" -pkg assets -tags "release" -ignore .DS_Store -o "./http/module/assets/bindata_release.go" ./public/fonts/... ./public/images/ ./public/scripts/... ./public/styles/... ./public/
