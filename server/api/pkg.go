package api

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"

	"github.com/KyleBanks/goggles/goggles"
)

// pkgList returns the names of each package in the $GOPATH.
func pkgList(w http.ResponseWriter, r *http.Request) {
	pkgs, err := goggles.ListPkgs()
	if err != nil {
		log.Fatal(err)
	}

	sort.Slice(pkgs, func(i, j int) bool {
		return pkgs[i].Name < pkgs[j].Name
	})
	json.NewEncoder(w).Encode(&pkgs)
}

// pkgDetails returns the full details of a package, identified by the
// parameter 'name'.
func pkgDetails(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	name := q.Get("name")

	p, err := goggles.Details(name)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(&p)
}
