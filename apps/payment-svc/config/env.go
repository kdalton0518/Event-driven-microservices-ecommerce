package config

import (
	"os"

	"github.com/spf13/viper"
)

var (
	GRPC_PORT    string
	DATABASE_URL string
	BROKER_URL   string
)

func LoadEnv() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		GRPC_PORT = os.Getenv("GRPC_PORT")
		DATABASE_URL = os.Getenv("DATABASE_URL")
		BROKER_URL = os.Getenv("BROKER_URL")
		return
	}

	GRPC_PORT = viper.GetString("GRPC_PORT")
	DATABASE_URL = viper.GetString("DATABASE_URL")
	BROKER_URL = viper.GetString("BROKER_URL")
}
