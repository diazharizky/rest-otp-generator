package health

import (
	"encoding/json"
	"net/http"

	httpUtils "github.com/diazharizky/rest-otp-generator/pkg/http"
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
		httpUtils.ResponseFatal(w, []string{err.Error()})
		return
	}
	httpUtils.ResponseSuccess(w, res)
}
