package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/KyleBanks/goggles/server"
	"github.com/alexflint/gallium"
)

const (
	port    = 10765
	logFile = "$HOME/Library/Logs/goggles.log"
	title   = ""
)

var window = gallium.WindowOptions{
	Shape: gallium.Rect{
		Width:  1200,
		Height: 800,
		Bottom: 400,
		Left:   400,
	},
	TitleBar:         true,
	Frame:            true,
	Resizable:        false,
	CloseButton:      true,
	MinButton:        true,
	FullScreenButton: false,
	Title:            title,
}

func main() {
	runtime.LockOSThread()
	gallium.RedirectStdoutStderr(os.ExpandEnv(logFile))
	gallium.Loop(os.Args, onReady)
}

func onReady(app *gallium.App) {
	w, err := app.OpenWindow(fmt.Sprintf("http://127.0.0.1:%v/static/index.html", port), window)
	if err != nil {
		log.Fatal(err)
	}

	server.Start(w, filepath.Dir(os.Args[0]), port)
}
