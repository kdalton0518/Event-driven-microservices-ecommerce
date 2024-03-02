package config

import (
	"os"

	"github.com/spf13/viper"
)

var (
	PORT         string
	DATABASE_URL string
	JWT_SECRET   string
)

func LoadEnv() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		PORT = os.Getenv("PORT")
		DATABASE_URL = os.Getenv("DATABASE_URL")
		return
	}

	PORT = viper.GetString("PORT")
	DATABASE_URL = viper.GetString("DATABASE_URL")
}
