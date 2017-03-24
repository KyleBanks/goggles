package server

import (
	"net/http"

	"github.com/KyleBanks/goggles/server/api"
	"github.com/KyleBanks/goggles/server/assets"
)

// New prepares and returns an HTTP ServeMux bound to the Goggles API and
// static assets.
func New(p api.Provider) *http.ServeMux {
	mux := http.NewServeMux()
	fs := http.FileServer(assets.FS())
	mux.Handle("/", http.StripPrefix("/static/", fs))
	api.Bind(p, mux)

	return mux
}
