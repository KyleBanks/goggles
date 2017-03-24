// Package main handles running Goggles as a standalone native application.
package main

import (
	"runtime"

	"github.com/KyleBanks/goggles/cmd"
)

const (
	title       = "Goggles"
	titleAbout  = "About"
	titleThanks = "Thanks"
	titleDebug  = "Debug"
	titleQuit   = "Quit"
)

func init() {
	runtime.LockOSThread()
}

func startServer() {
	cmd.StartServer()
}
