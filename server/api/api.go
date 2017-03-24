package api

import (
	"fmt"
	"io"
	"net/http"

	"github.com/KyleBanks/goggles/conf"
	"github.com/KyleBanks/goggles/resolver"
)

var provider Provider

// Provider is a type that provides access to package data, the host operating system,
// and anything else the API requires to function.
type Provider interface {
	List() ([]*resolver.Package, error)
	Details(string) (*resolver.Package, error)

	OpenFileExplorer(string)
	OpenTerminal(string)
	OpenBrowser(string)

	Preferences() *conf.Config
	UpdatePreferences(*conf.Config)
}

// Bind attaches the API routes to the default HTTP server.
func Bind(p Provider, mux *http.ServeMux) {
	provider = p

	// PKGs
	mux.HandleFunc("/api/pkg/list", pkgList)
	mux.HandleFunc("/api/pkg/details", pkgDetails)

	// Applications
	mux.HandleFunc("/api/open/file-explorer", openFileExplorer)
	mux.HandleFunc("/api/open/terminal", openTerminal)
	mux.HandleFunc("/api/open/url", openURL)

	// Preferences
	mux.HandleFunc("/api/preferences/", getPreferences)
	mux.HandleFunc("/api/preferences/update", updatePreferences)
}

// outputEmpty prints an empty JSON response to a Writer.
func outputEmpty(w io.Writer) {
	fmt.Fprintf(w, "{}")
}
