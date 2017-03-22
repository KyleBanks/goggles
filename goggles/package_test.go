package goggles

import (
	"testing"
)

func Test_NewPackage(t *testing.T) {
	path := "github.com/KyleBanks/goggles/goggles"
	p, err := NewPackage(path)
	if err != nil {
		t.Fatal(err)
	}

	if p.Name != path {
		t.Fatalf("Unexpected Name, expected=%v, got=%v", path, p.Name)
	}
}
