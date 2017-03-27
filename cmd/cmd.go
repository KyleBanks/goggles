// Package cmd provides shared functionality for the multiple run-modes.
package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/KyleBanks/goggles/pkg/release"
	"github.com/KyleBanks/goggles/pkg/sys"
	"github.com/KyleBanks/goggles/resolver"
	"github.com/KyleBanks/goggles/server"
	"github.com/KyleBanks/goggles/server/api"
)

const (
	owner = "KyleBanks"
	repo  = "goggles"
)

var version string

var (
	// port is the port number to listen on.
	port = 10765

	// Index is the URL of the root index.html file.
	Index = fmt.Sprintf("http://127.0.0.1:%v/static/index.html", port)

	defaultProvider api.Provider = provider{resolver.Resolver{}}

	aboutURL  = fmt.Sprintf("https://github.com/%v/%v", owner, repo)
	thanksURL = fmt.Sprintf("https://github.com/%v/%v#thanks", owner, repo)
)

func init() {
	log.Printf("v%v", version)

	initConfig()
	log.Printf("$GOPATH=%v, srcdir=%v", sys.Gopath(), sys.Srcdir())

	go checkRelease()
}

func initConfig() {
	c := defaultProvider.Preferences()

	// Update the $GOPATH if a custom value is set.
	if len(c.Gopath) > 0 {
		sys.SetGopath(c.Gopath)
	}
}

// StartServer starts the application server.
func StartServer() {
	api := server.New(defaultProvider)
	addr := fmt.Sprintf(":%v", port)

	log.Fatal(http.ListenAndServe(addr, api))
}

// OpenAbout opens the 'About' page in a web browser.
func OpenAbout() {
	defaultProvider.OpenBrowser(aboutURL)
}

// OpenThanks opens the 'Thanks' page in a web browser.
func OpenThanks() {
	defaultProvider.OpenBrowser(thanksURL)
}

// Quit terminates the running application.
func Quit() {
	os.Exit(0)
}

// checkRelease checks if there is a new release of Goggles available, and prompts
// the user to update if necessary.
func checkRelease() {
	latest, version, err := release.IsLatest(owner, repo, version)
	if err != nil {
		log.Printf("error checking for latest release: %v", err)
		return
	} else if latest {
		return
	}

	log.Printf("\n\n***\nVersion %v available, to update now:\ngo get -u github.com/%v/%v/cmd/goggles\n***\n\n", owner, repo, version)
}
