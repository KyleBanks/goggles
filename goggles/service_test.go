package goggles

import (
	"testing"
)

func TestService_List(t *testing.T) {
	var s Service
	pkgs, err := s.List()
	if err != nil {
		t.Fatal(err)
	}

	if len(pkgs) == 0 {
		t.Fatal("Expected non-empty pkgs")
	}

	var found bool
	for _, p := range pkgs {
		if p.Name == "github.com/KyleBanks/goggles/goggles" {
			found = true
		}

		if p.Name == "" {
			t.Fatalf("Unexected empty name: %v", p)
		} else if p.Docs != nil {
			t.Fatalf("Unexpected docs, expected=%v, got=%v", nil, p.Docs)
		}
	}

	if !found {
		t.Fatal("Expected goggles pkg to be returned")
	}
}

func TestService_Details(t *testing.T) {
	var s Service
	p, err := s.Details("github.com/KyleBanks/goggles/goggles")
	if err != nil {
		t.Fatal(err)
	}

	if p.Docs.Type != PackageDoc {
		t.Fatalf("Unexpected Type, expected=%v, got=%v", PackageDoc, p.Docs.Type)
	} else if name := "goggles"; p.Docs.Name != name {
		t.Fatalf("Unexpected Name, expected=%v, got=%v", name, p.Docs.Name)
	} else if imp := "import \"github.com/KyleBanks/goggles/goggles\""; p.Docs.Import != imp {
		t.Fatalf("Unexpected Import, expected=%v, got=%v", imp, p.Docs.Import)
	} else if repo := "github.com/KyleBanks/goggles"; p.Docs.Repository != repo {
		t.Fatalf("Unexpected Repository, expected=%v, got=%v", repo, p.Docs.Repository)
	}
}
