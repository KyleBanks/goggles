package api

import (
	"net/http"
)

// debug opens the developer tools for debugging.
func debug(w http.ResponseWriter, r *http.Request) {
	provider.OpenDevTools()
	outputEmpty(w)
}
