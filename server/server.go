package server

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

var devTools DevTooler

// DevTooler defines a type that can be used to display developer tools.
type DevTooler interface {
	OpenDevTools()
}

// Start prepares and starts the HTTP server.
//
// The root parameter should be to the parent of the running binary where assets can
// be found. For example, in the following case "/foo/bar" would be the root.
//
// /foo/bar
//    /goggles
//    /static/...
func Start(d DevTooler, root string, port int) {
	log.Printf("server.Start(%v, %v)", root, port)
	root = filepath.Join(root, "static")

	fs := http.FileServer(http.Dir(root))
	http.Handle("/", http.StripPrefix("/static/", fs))
	bindAPIRoutes()

	http.ListenAndServe(fmt.Sprintf(":%v", port), wrap(http.DefaultServeMux))
}

// wrap wraps the Handler provided with additional functionality, such as logging.
func wrap(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
