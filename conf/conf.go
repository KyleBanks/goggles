// Package conf provides Goggles configuration.
package conf

import (
	"github.com/KyleBanks/go-kit/storage"
)

type saveLoader interface {
	Save(interface{}) error
	Load(interface{}) error
}

var store saveLoader = storage.NewFileStore("goggles", "goggles.conf")

// Config contains Goggles configuration.
type Config struct {
	Gopath string `json:"gopath"`
}

// Get returns the persisted Config.
func Get() *Config {
	var c Config
	store.Load(&c)
	return &c
}

// Save persists the provided Config.
func Save(c *Config) error {
	return store.Save(c)
}
