package api

import (
	"encoding/json"
	"net/http"

	"github.com/KyleBanks/goggles/conf"
)

func getPreferences(w http.ResponseWriter, r *http.Request) {
	p := provider.Preferences()

	json.NewEncoder(w).Encode(&p)
}

func updatePreferences(w http.ResponseWriter, r *http.Request) {
	var c conf.Config
	c.Gopath = r.URL.Query().Get("gopath")

	provider.UpdatePreferences(&c)

	outputEmpty(w)
}
