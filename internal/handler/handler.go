package handler

import (
	"github.com/go-chi/chi"
)

func Router() chi.Router {
	r := chi.NewRouter()
	r.Mount("/health", healthHandler())
	r.Mount("/otp", otpHandler())
	return r
}
