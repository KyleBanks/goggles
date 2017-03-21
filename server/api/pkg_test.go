package api

import (
	"net/http/httptest"
	"testing"

	"github.com/KyleBanks/goggles/goggles"
)

func Test_pkgList(t *testing.T) {
	m := setup()

	expect := []*goggles.Package{
		&goggles.Package{Docs: &goggles.Doc{Name: "Name 1"}},
		&goggles.Package{Docs: &goggles.Doc{Name: "Name 2"}},
		&goggles.Package{Docs: &goggles.Doc{Name: "Name 3"}},
	}
	m.ListFn = func() ([]*goggles.Package, error) {
		return expect, nil
	}

	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	pkgList(w, r)

	var pkgs []*goggles.Package
	validateResponse(t, w, &pkgs)

	if len(pkgs) != len(expect) {
		t.Fatalf("Unexpected number of pkgs, expected=%v, got=%v", len(expect), len(pkgs))
	}
	for idx, p := range pkgs {
		if p.Docs.Name != expect[idx].Docs.Name {
			t.Fatalf("[%v] Unexpected pkg, expected=%v, got=%v", idx, expect[idx].Name, p.Name)
		}
	}
}

func Test_pkgDetails(t *testing.T) {
	m := setup()

	name := "foo/bar"
	expect := goggles.Package{
		Docs: &goggles.Doc{
			Name:   name,
			Import: "import \"foo/bar\"",
		},
	}
	m.DetailsFn = func(s string) (*goggles.Package, error) {
		if s != name {
			t.Fatalf("Unexpected name, expected=%v, got=%v", name, s)
		}
		return &expect, nil
	}

	r := httptest.NewRequest("GET", "/?name="+name, nil)
	w := httptest.NewRecorder()
	pkgDetails(w, r)

	var pkg goggles.Package
	validateResponse(t, w, &pkg)

	if pkg.Docs.Name != expect.Docs.Name || pkg.Docs.Import != expect.Docs.Import {
		t.Fatalf("Unexpected package returned, expected=%v, got=%v", expect, pkg)
	}
}
