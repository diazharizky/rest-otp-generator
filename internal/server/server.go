package server

import (
	"fmt"
	"net/http"

	"github.com/diazharizky/rest-otp-generator/configs"
	"github.com/diazharizky/rest-otp-generator/internal/handler"
	log "github.com/sirupsen/logrus"
)

func Serve() {
	listenHost := configs.Cfg.GetString("listen.host")
	listenPort := configs.Cfg.GetInt("listen.port")
	log.Info(fmt.Sprintf("Listening on %s:%d!", listenHost, listenPort))
	http.ListenAndServe(fmt.Sprintf(":%d", listenPort), handler.Router())
}
