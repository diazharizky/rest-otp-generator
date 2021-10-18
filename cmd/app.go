package cmd

import (
	"github.com/diazharizky/rest-otp-generator/internal/server"
	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command

func init() {
	rootCmd = &cobra.Command{
		Use:   "app",
		Short: "Start OTP generator app",
		Run: func(cmd *cobra.Command, args []string) {
			server.Serve()
		},
	}
}

func Execute() error {
	return rootCmd.Execute()
}
