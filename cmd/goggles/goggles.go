package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/KyleBanks/goggles/conf"
	"github.com/KyleBanks/goggles/goggles"
	"github.com/KyleBanks/goggles/pkg/sys"
	"github.com/KyleBanks/goggles/server"
)

const (
	port = 10765

	title       = "Goggles"
	titleAbout  = "About"
	titleThanks = "Thanks"
	titleDebug  = "Debug"
	titleQuit   = "Quit"

	aboutURL  = "https://github.com/KyleBanks/goggles"
	thanksURL = "https://github.com/KyleBanks/goggles#thanks"
)

var (
	logFile = os.ExpandEnv("$HOME/Library/Logs/goggles.log")
	index   = fmt.Sprintf("http://127.0.0.1:%v/static/index.html", port)
)

func init() {
	runtime.LockOSThread()

	// Update the $GOPATH if a custom value is set.
	c := conf.Get()
	if c != nil && len(c.Gopath) > 0 {
		sys.SetGopath(c.Gopath)
	}
}

func startServer() {
	log.Printf("$GOPATH=%v, srcdir=%v", sys.Gopath(), sys.Srcdir())

	p := provider{goggles.Service{}}
	api := server.New(p, filepath.Dir(os.Args[0]))
	addr := fmt.Sprintf(":%v", port)
	log.Fatal(http.ListenAndServe(addr, api))
}

func openAbout() {
	sys.OpenBrowser(aboutURL)
}

func openThanks() {
	sys.OpenBrowser(thanksURL)
}

func quit() {
	os.Exit(0)
}
