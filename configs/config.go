package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("listen.host", "0.0.0.0")
	viper.SetDefault("listen.port", 3000)
}

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("cannot read config file: %w", err))
	}
}
