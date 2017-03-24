// Package cmd provides shared functionality for the multiple run-modes.
package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/KyleBanks/goggles"
	"github.com/KyleBanks/goggles/conf"
	"github.com/KyleBanks/goggles/pkg/sys"
	"github.com/KyleBanks/goggles/server"
)

const (
	aboutURL  = "https://github.com/KyleBanks/goggles"
	thanksURL = "https://github.com/KyleBanks/goggles#thanks"
)

var (
	// port is the port number to listen on.
	port = 10765

	// Index is the URL of the root index.html file.
	Index = fmt.Sprintf("http://127.0.0.1:%v/static/index.html", port)
)

func init() {
	// Update the $GOPATH if a custom value is set.
	c := conf.Get()
	if c != nil && len(c.Gopath) > 0 {
		sys.SetGopath(c.Gopath)
	}

	log.Printf("$GOPATH=%v, srcdir=%v", sys.Gopath(), sys.Srcdir())
}

// StartServer starts the application server.
func StartServer() {
	p := provider{goggles.Service{}}
	api := server.New(p)
	addr := fmt.Sprintf(":%v", port)

	log.Fatal(http.ListenAndServe(addr, api))
}

// OpenAbout opens the 'About' page in a web browser.
func OpenAbout() {
	sys.OpenBrowser(aboutURL)
}

// OpenThanks opens the 'Thanks' page in a web browser.
func OpenThanks() {
	sys.OpenBrowser(thanksURL)
}

// Quit terminates the running application.
func Quit() {
	os.Exit(0)
}
