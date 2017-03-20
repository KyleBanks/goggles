package api

import (
	"net/http"

	"github.com/KyleBanks/goggles/pkg"
)

var (
	provider Provider
)

// Provider is a type that provides access to package data, the host operating system,
// and anything else the API requires to function.
type Provider interface {
	List() ([]*pkg.Package, error)
	Details(string) (*pkg.Package, error)

	OpenDevTools()
	OpenFileExplorer(string)
}

// Bind attaches the API routes to the default HTTP server.
func Bind(p Provider) {
	provider = p

	// PKGs
	http.HandleFunc("/api/pkg/list", pkgList)
	http.HandleFunc("/api/pkg/details", pkgDetails)

	// Applications
	http.HandleFunc("/api/open/file-explorer", openFileExplorer)

	// Misc.
	http.HandleFunc("/api/debug", debug)
}
