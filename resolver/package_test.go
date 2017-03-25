package resolver

import (
	"testing"
)

func Test_NewPackage(t *testing.T) {
	path := "github.com/KyleBanks/goggles/resolver"
	p, err := NewPackage(path)
	if err != nil {
		t.Fatal(err)
	}

	if p.Name != path {
		t.Fatalf("Unexpected Name, expected=%v, got=%v", path, p.Name)
	}
}
