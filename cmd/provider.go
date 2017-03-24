package cmd

import (
	"log"

	"github.com/KyleBanks/goggles/conf"
	"github.com/KyleBanks/goggles/pkg/sys"
	"github.com/KyleBanks/goggles/resolver"
)

// provider wraps the Goggles packages into a single type
// that can provide all functionality to the API.
type provider struct {
	resolver.Resolver
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

func (provider) Preferences() *conf.Config {
	c := conf.Get()

	// Set defaults
	if len(c.Gopath) == 0 {
		c.Gopath = sys.RawGopath()
	}

	return c
}

func (provider) UpdatePreferences(c *conf.Config) {
	if err := conf.Save(c); err != nil {
		log.Printf("Failed to save Config [%v] due to: %v", c, err)
	}

	sys.SetGopath(c.Gopath)
}
