package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

var Config *viper.Viper

func init() {
	Config = viper.New()
	Config.SetDefault("listen.host", "0.0.0.0")
	Config.SetDefault("listen.port", 5000)

	loadConfig()
}

func loadConfig() {
	Config.SetConfigName("config")
	Config.SetConfigType("yaml")
	Config.AddConfigPath("./configs")
	if err := Config.ReadInConfig(); err != nil {
		panic(fmt.Errorf("cannot read config file: %w", err))
	}
}
