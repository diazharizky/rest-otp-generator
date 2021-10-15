package routes

import (
	"github.com/go-chi/chi"

	"github.com/diazharizky/rest-otp-generator/internal/health"
	"github.com/diazharizky/rest-otp-generator/internal/otp"
)

func Router() chi.Router {
	r := chi.NewRouter()
	r.Mount("/otp", otp.Handler())
	r.Mount("/health", health.Handler())
	return r
}
