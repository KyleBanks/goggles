package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"

	"github.com/KyleBanks/goggles/goggles"
	"github.com/alexflint/gallium"
)

type API struct {
	w *gallium.Window
}

// bind attaches the API's routes to the default HTTP server.
func (a *API) bind() {
	http.HandleFunc("/api/pkg/list", a.pkgList)
	http.HandleFunc("/api/devtools", a.showDevTools)
}

// pkgList returns the names of each package in the $GOPATH.
func (*API) pkgList(w http.ResponseWriter, r *http.Request) {
	pkgs, err := goggles.ListPkgs()
	if err != nil {
		log.Fatal(err)
	}

	sort.Slice(pkgs, func(i, j int) bool {
		return pkgs[i].Name < pkgs[j].Name
	})
	json.NewEncoder(w).Encode(&pkgs)
}

func (a *API) showDevTools(w http.ResponseWriter, r *http.Request) {
	a.w.OpenDevTools()
	fmt.Fprintf(w, "{}")
}
