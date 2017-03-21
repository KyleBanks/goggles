package api

import (
	"fmt"
	"io"
	"net/http"

	"github.com/KyleBanks/goggles/goggles"
)

var (
	provider Provider
)

// Provider is a type that provides access to package data, the host operating system,
// and anything else the API requires to function.
type Provider interface {
	List() ([]*goggles.Package, error)
	Details(string) (*goggles.Package, error)

	OpenDevTools()
	OpenFileExplorer(string)
}

// Bind attaches the API routes to the default HTTP server.
func Bind(p Provider, mux *http.ServeMux) {
	provider = p

	// PKGs
	mux.HandleFunc("/api/pkg/list", pkgList)
	mux.HandleFunc("/api/pkg/details", pkgDetails)

	// Applications
	mux.HandleFunc("/api/open/file-explorer", openFileExplorer)

	// Misc.
	mux.HandleFunc("/api/debug", debug)
}

func outputEmpty(w io.Writer) {
	fmt.Fprintf(w, "{}")
}
