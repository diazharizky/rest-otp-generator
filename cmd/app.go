package cmd

import (
	"fmt"
	"net/http"

	"github.com/diazharizky/rest-otp-generator/configs"
	"github.com/diazharizky/rest-otp-generator/internal/routes"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command

func init() {
	rootCmd = &cobra.Command{
		Use:   "app",
		Short: "Start main app",
		Run: func(cmd *cobra.Command, args []string) {
			serve()
		},
	}
}

func serve() {
	listenPort := configs.Config.GetInt("listen.port")
	log.Info("Listening on " + configs.Config.GetString("listen.host") + ":" + fmt.Sprintf("%d", listenPort) + "!")
	http.ListenAndServe(fmt.Sprintf(":%d", listenPort), routes.Router())
}

func Execute() error {
	return rootCmd.Execute()
}
