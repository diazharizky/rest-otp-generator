package server

import (
	"fmt"
	"net/http"

	"github.com/diazharizky/rest-otp-generator/configs"
	"github.com/diazharizky/rest-otp-generator/internal/health"
	"github.com/diazharizky/rest-otp-generator/internal/otp"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func router() chi.Router {
	r := chi.NewRouter()
	r.Mount("/health", health.Handler())
	r.Mount("/otp", otp.Handler())

	return r
}

func Serve() {
	listenHost := configs.Cfg.GetString("listen.host")
	listenPort := configs.Cfg.GetInt("listen.port")
	log.Info(fmt.Sprintf("Listening on %s:%d!", listenHost, listenPort))
	http.ListenAndServe(fmt.Sprintf(":%d", listenPort), router())
}
