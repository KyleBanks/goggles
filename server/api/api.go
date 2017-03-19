package api

import (
	"net/http"

	"github.com/KyleBanks/goggles/goggles"
)

var (
	devTools DevTooler
	packager Packager
)

// DevTooler defines a type that can be used to display developer tools.
type DevTooler interface {
	OpenDevTools()
}

// Packager is a type that provides access to package data.
type Packager interface {
	List() ([]*goggles.Pkg, error)
	Details(string) (*goggles.Pkg, error)
	OpenFileExplorer(string)
}

// Bind attaches the API routes to the default HTTP server.
func Bind(d DevTooler, p Packager) {
	devTools = d
	packager = p

	// PKGs
	http.HandleFunc("/api/pkg/list", pkgList)
	http.HandleFunc("/api/pkg/details", pkgDetails)

	// Applications
	http.HandleFunc("/api/open/file-explorer", openFileExplorer)

	// Misc.
	http.HandleFunc("/api/debug", debug)
}
