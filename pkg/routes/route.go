package routes

import (
	"github.com/go-chi/chi"

	"github.com/diazharizky/rest-otp-generator/pkg/otp"
)

func Router() chi.Router {
	r := chi.NewRouter()
	r.Mount("/otp", otp.Handler())
	return r
}
