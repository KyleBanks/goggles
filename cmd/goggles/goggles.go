package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/KyleBanks/goggles/goggles"
	"github.com/KyleBanks/goggles/server"
	"github.com/alexflint/gallium"
)

var (
	port    = 10765
	logFile = os.ExpandEnv("$HOME/Library/Logs/goggles.log")
	index   = fmt.Sprintf("http://127.0.0.1:%v/static/index.html", port)

	window = gallium.WindowOptions{
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
		Title:            "",
	}
)

func init() {
	runtime.LockOSThread()
}

func main() {
	gallium.RedirectStdoutStderr(logFile)

	log.Fatal(gallium.Loop(os.Args, onReady))
}

func onReady(app *gallium.App) {
	w, err := app.OpenWindow(index, window)
	if err != nil {
		log.Fatal(err)
	}

	addr := fmt.Sprintf(":%v", port)
	api := server.New(provider{
		w,
		goggles.Default,
	}, filepath.Dir(os.Args[0]))

	log.Fatal(http.ListenAndServe(addr, api))
}
