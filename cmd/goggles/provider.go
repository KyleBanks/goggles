package main

import (
	"github.com/KyleBanks/goggles/goggles"
	"github.com/KyleBanks/goggles/pkg/sys"
	"github.com/alexflint/gallium"
)

type provider struct {
	*gallium.Window
	goggles.Service
}

func (provider) OpenFileExplorer(n string) {
	sys.OpenFileExplorer(n)
}
