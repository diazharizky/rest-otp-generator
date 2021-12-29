package handler

import (
	"encoding/json"
	"net/http"

	core "github.com/diazharizky/rest-otp-generator/internal/core"
	"github.com/diazharizky/rest-otp-generator/pkg/health"
	httpUtils "github.com/diazharizky/rest-otp-generator/pkg/http"
	"github.com/go-chi/chi"
)

type healthManager interface {
	HealthCheck() health.HealthStatus
}

type healthDriver struct {
	healthManager
}

var hdr healthDriver

func init() {
	hdr = healthDriver{&core.Core}
}

func healthHandler() (r *chi.Mux) {
	r = chi.NewRouter()
	r.Get("/", healthCheck)
	return
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	res, err := json.Marshal(hdr.HealthCheck())
	if err != nil {
		httpUtils.ResponseFatal(w, []string{err.Error()})
		return
	}
	httpUtils.ResponseSuccess(w, res)
}
