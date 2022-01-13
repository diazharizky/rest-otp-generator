package interfaces

import (
	"net/http"

	"github.com/diazharizky/rest-otp-generator/internal/application"
	httpUtils "github.com/diazharizky/rest-otp-generator/pkg/http"
	"github.com/go-chi/chi"
)

type healthHandler struct {
	ha application.HealthAppInterface
}

func NewHealthHandler(ha application.HealthAppInterface) healthHandler {
	return healthHandler{ha}
}

func (h *healthHandler) getHandler() (r *chi.Mux) {
	r = chi.NewRouter()
	r.Get("/", h.healthCheck)
	return
}

func (h *healthHandler) healthCheck(w http.ResponseWriter, r *http.Request) {
	httpUtils.ResponseSuccess(w, h.ha.HealthCheck())
}
