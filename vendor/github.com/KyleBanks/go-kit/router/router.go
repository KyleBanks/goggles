// Package router defines the Route interface, and registers routes to an http server.
package router

import (
	"fmt"
	"net/http"

	"github.com/KyleBanks/go-kit/timer"
)

// Handler defines a function that accepts an HTTP request and returns a Response.
type Handler func(http.ResponseWriter, *http.Request)

// Server defines an interface for the provided server to comply with.
type Server interface {
	HandleFunc(string, func(http.ResponseWriter, *http.Request))
}

// Route defines a URL path and function to execute when the URL
// is accessed.
type Route struct {
	Path   string // The URL path to listen for (i.e. "/api")
	Handle Handler
}

// Register registers each Route with the Server provided.
//
// Each Route will be wrapped in a middleware function that adds trace logging.
func Register(s Server, routes []Route) {
	for _, route := range routes {
		s.HandleFunc(route.Path, handleWrapper(route))
	}
}

// handleWrapper returns a request handling function that wraps the provided route.
func handleWrapper(route Route) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		// Enable CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		t := timer.NewLogger(fmt.Sprintf("%v: {Query: %v, Form: %v}", r.URL.Path, r.URL.RawQuery, r.PostForm))
		defer t()

		route.Handle(w, r)
	}
}

// Param returns a POST/GET parameter from the request.
//
// If the parameter is found in the POST and the GET parameter set, the POST parameter
// will be given priority.
func Param(r *http.Request, key string) string {
	r.ParseForm()

	val := r.PostForm.Get(key)
	if len(val) != 0 {
		return val
	}

	return r.URL.Query().Get(key)
}

// HasParam returns a boolean indicating if the request has a particular parameter.
func HasParam(r *http.Request, key string) bool {
	return len(Param(r, key)) > 0
}
