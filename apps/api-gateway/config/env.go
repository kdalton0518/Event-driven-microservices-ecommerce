package config

import (
	"os"

	"github.com/spf13/viper"
)

var (
	PORT string
)

func LoadEnv() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		PORT = os.Getenv("PORT")
		return
	}

	PORT = viper.GetString("PORT")
}
