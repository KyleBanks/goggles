package api

import (
	"net/http"
)

var devTools DevTooler

// DevTooler defines a type that can be used to display developer tools.
type DevTooler interface {
	OpenDevTools()
}

// Bind attaches the API routes to the default HTTP server.
func Bind(d DevTooler) {
	devTools = d

	// PKGs
	http.HandleFunc("/api/pkg/list", pkgList)
	http.HandleFunc("/api/pkg/details", pkgDetails)

	// Misc.
	http.HandleFunc("/api/debug", debug)
}
