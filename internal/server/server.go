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
	listenPort := configs.Cfg.GetInt("listen.port")
	log.Info("Listening on " + configs.Cfg.GetString("listen.host") + ":" + fmt.Sprintf("%d", listenPort) + "!")
	http.ListenAndServe(fmt.Sprintf(":%d", listenPort), router())
}
