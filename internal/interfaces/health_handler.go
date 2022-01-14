package interfaces

import (
	"net/http"

	"github.com/diazharizky/rest-otp-generator/internal/application"
	httpUtils "github.com/diazharizky/rest-otp-generator/pkg/http"
	"github.com/go-chi/chi"
)

type HealthHandler struct {
	ha application.HealthAppInterface
}

func NewHealthHandler(ha application.HealthAppInterface) HealthHandler {
	return HealthHandler{ha}
}

func (h *HealthHandler) getHandler() (r *chi.Mux) {
	r = chi.NewRouter()
	r.Get("/", h.HealthCheck)
	return
}

func (h *HealthHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	httpUtils.ResponseSuccess(w, h.ha.HealthCheck())
}
