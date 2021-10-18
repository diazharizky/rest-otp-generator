package health

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

var mCore core

func Handler() (r *chi.Mux) {
	r = chi.NewRouter()
	r.Get("/", healthCheck)
	return
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	status := mCore.health().DB
	res := map[string]interface{}{"status": status}
	resByte, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resByte)
}