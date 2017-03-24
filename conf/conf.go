// Package conf provides Goggles configuration.
package conf

import (
	"github.com/KyleBanks/go-kit/storage"
	"github.com/KyleBanks/goggles/pkg/sys"
)

type saveLoader interface {
	Save(interface{}) error
	Load(interface{}) error
}

var store saveLoader = storage.NewFileStore("goggles", "goggles.conf")

// Config contains Goggles configuration.
type Config struct {
	// user configurable

	Gopath string `json:"gopath"`

	// system values
	// these values will be ignored in the persisted file

	CanOpenTerminal bool `json:"canOpenTerminal"`
}

// Get returns the persisted Config.
func Get() *Config {
	var c Config
	store.Load(&c)

	// Apply system-configurations
	c.CanOpenTerminal = sys.CanOpenTerminal

	return &c
}

// Save persists the provided Config.
func Save(c *Config) error {
	return store.Save(c)
}
