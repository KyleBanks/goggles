package main

import (
	"github.com/KyleBanks/goggles/goggles"
	"github.com/KyleBanks/goggles/pkg/sys"
)

type provider struct {
	goggles.Service
}

func (provider) OpenFileExplorer(n string) {
	sys.OpenFileExplorer(n)
}

func (provider) OpenTerminal(n string) {
	sys.OpenTerminal(n)
}

func (provider) OpenBrowser(n string) {
	sys.OpenBrowser(n)
}
