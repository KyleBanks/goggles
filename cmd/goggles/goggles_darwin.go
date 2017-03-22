// +build darwin

package main

import (
	"log"
	"os"

	"github.com/alexflint/gallium"
)

var (
	window *gallium.Window

	opts = gallium.WindowOptions{
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

	menu = []gallium.Menu{
		gallium.Menu{
			Title: title,
			Entries: []gallium.MenuEntry{
				gallium.MenuItem{
					Title:   titleAbout,
					OnClick: openAbout,
				},
				gallium.MenuItem{
					Title:   titleThanks,
					OnClick: openThanks,
				},

				gallium.Separator,

				gallium.MenuItem{
					Title:    titleDebug,
					Shortcut: gallium.MustParseKeys("CMD+d"),
					OnClick: func() {
						window.OpenDevTools()
					},
				},
				gallium.MenuItem{
					Title:    titleQuit,
					Shortcut: gallium.MustParseKeys("CMD+q"),
					OnClick:  quit,
				},
			},
		},
	}
)

func main() {
	gallium.RedirectStdoutStderr(logFile)
	log.Fatal(gallium.Loop(os.Args, onReady))
}

func onReady(app *gallium.App) {
	var err error
	window, err = app.OpenWindow(index, opts)
	if err != nil {
		log.Fatal(err)
	}
	app.SetMenu(menu)

	startServer()
}
