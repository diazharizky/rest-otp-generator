package server

import (
	"fmt"
	"net/http"

	"github.com/diazharizky/rest-otp-generator/configs"
	"github.com/diazharizky/rest-otp-generator/internal/routes"
	log "github.com/sirupsen/logrus"
)

func Serve() {
	listenPort := configs.Cfg.GetInt("listen.port")
	log.Info("Listening on " + configs.Cfg.GetString("listen.host") + ":" + fmt.Sprintf("%d", listenPort) + "!")
	http.ListenAndServe(fmt.Sprintf(":%d", listenPort), routes.Router())
}
