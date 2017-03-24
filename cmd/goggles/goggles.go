// Package main handles running Goggles in a web browser.
package main

import (
	"time"

	"github.com/KyleBanks/goggles/cmd"
	"github.com/KyleBanks/goggles/pkg/sys"
)

func main() {
	go func() {
		time.Sleep(time.Millisecond * 500)
		sys.OpenBrowser(cmd.Index)
	}()

	cmd.StartServer()
}
