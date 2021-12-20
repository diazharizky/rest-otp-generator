package configs

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var Cfg *viper.Viper

func init() {
	Cfg = viper.New()
	Cfg.SetDefault("listen.host", "0.0.0.0")
	Cfg.SetDefault("listen.port", 8080)

	loadConfig()
}

func loadConfig() {
	Cfg.SetConfigName("default")
	Cfg.SetConfigType("yaml")

	configFilePath := os.Getenv("CONFIG_FILE_PATH")
	if len(configFilePath) <= 0 {
		configFilePath = "./configs"
	}

	Cfg.AddConfigPath(configFilePath)
	if err := Cfg.ReadInConfig(); err != nil {
		panic(fmt.Errorf("cannot read config file: %w", err))
	}
}
