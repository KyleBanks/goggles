package api

import (
	"net/http"
)

// openFileExplorer opens the system File Explorer to the package provided.
func openFileExplorer(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	name := q.Get("name")

	provider.OpenFileExplorer(name)
	outputEmpty(w)
}

// openTerminal opens the system terminal to the package provided.
func openTerminal(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	name := q.Get("name")

	provider.OpenTerminal(name)
	outputEmpty(w)
}

// openURL opens the default browser to the specified URL.
func openURL(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	url := q.Get("url")

	provider.OpenBrowser(url)
	outputEmpty(w)
}
