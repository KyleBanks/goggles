package server

import (
	"net/http"
	"path/filepath"

	"github.com/KyleBanks/goggles/server/api"
)

// New prepares and returns an HTTP ServeMux bound to the Goggles API and
// static assets.
//
// The root parameter should be to the parent of the running binary where assets can
// be found. For example, in the following case "/foo/bar" would be the root.
//
// /foo/bar
//    /goggles
//    /static/...
func New(p api.Provider, root string) *http.ServeMux {
	mux := http.NewServeMux()
	dir := http.Dir(filepath.Join(root, "static"))
	fs := http.FileServer(dir)
	mux.Handle("/", http.StripPrefix("/static/", fs))
	api.Bind(p, mux)

	return mux
}
