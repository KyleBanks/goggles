package api

import (
	"fmt"
	"net/http"
)

// debug opens the developer tools for debugging.
func debug(w http.ResponseWriter, r *http.Request) {
	provider.OpenDevTools()
	fmt.Fprintf(w, "{}")
}
