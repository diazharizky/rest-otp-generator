package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

var Cfg *viper.Viper

func init() {
	Cfg = viper.New()
	Cfg.SetDefault("listen.host", "0.0.0.0")
	Cfg.SetDefault("listen.port", 5000)

	loadConfig()
}

func loadConfig() {
	Cfg.SetConfigName("config")
	Cfg.SetConfigType("yaml")
	Cfg.AddConfigPath("./configs")
	if err := Cfg.ReadInConfig(); err != nil {
		panic(fmt.Errorf("cannot read config file: %w", err))
	}
}
