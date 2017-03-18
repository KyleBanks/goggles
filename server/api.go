package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"

	"github.com/KyleBanks/goggles/goggles"
)

// bindAPIRoutes attaches the API routes to the default HTTP server.
func bindAPIRoutes() {
	http.HandleFunc("/api/debug", debug)
	http.HandleFunc("/api/pkg/list", pkgList)
}

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

func debug(w http.ResponseWriter, r *http.Request) {
	if devTools != nil {
		devTools.OpenDevTools()
	}
	fmt.Fprintf(w, "{}")
}
