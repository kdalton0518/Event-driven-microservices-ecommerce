package config

import (
	"os"

	"github.com/spf13/viper"
)

var (
	PORT         string
	GRPC_PORT    string
	DATABASE_URL string
)

func LoadEnv() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		PORT = os.Getenv("PORT")
		GRPC_PORT = os.Getenv("GRPC_PORT")
		DATABASE_URL = os.Getenv("DATABASE_URL")
		return
	}

	PORT = viper.GetString("PORT")
	GRPC_PORT = viper.GetString("GRPC_PORT")
	DATABASE_URL = viper.GetString("DATABASE_URL")
}
