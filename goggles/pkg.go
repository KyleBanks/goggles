package goggles

import (
	"github.com/KyleBanks/depth"
)

// Pkg represents a go source package.
type Pkg struct {
	depth.Pkg

	Docs string `json:"docs"`
}
