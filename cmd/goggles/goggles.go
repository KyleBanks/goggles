// Package main handles running Goggles in a web browser.
package main

import (
	"time"

	"github.com/KyleBanks/goggles/cmd"
	"github.com/KyleBanks/goggles/pkg/sys"
)

func main() {
	go openBrowser(time.Millisecond * 500)
	cmd.StartServer()
}

func openBrowser(delay time.Duration) {
	time.Sleep(delay)
	sys.OpenBrowser(cmd.Index)
}
