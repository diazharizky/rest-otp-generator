package health

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

func Handler() (r *chi.Mux) {
	r = chi.NewRouter()
	r.Get("/", health)

	return
}

func health(w http.ResponseWriter, r *http.Request) {
	res, err := json.Marshal(c.healthCheck())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
