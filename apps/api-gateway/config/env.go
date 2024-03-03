package config

import (
	"os"

	"github.com/spf13/viper"
)

var (
	PORT                   string
	GRPC_HOST_PRODUCT_SVC  string
	GRPC_HOST_CUSTOMER_SVC string
)

func LoadEnv() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		PORT = os.Getenv("PORT")
		GRPC_HOST_PRODUCT_SVC = os.Getenv("GRPC_HOST_PRODUCT_SVC")
		GRPC_HOST_CUSTOMER_SVC = os.Getenv("GRPC_HOST_CUSTOMER_SVC")
		return
	}

	PORT = viper.GetString("PORT")
	GRPC_HOST_PRODUCT_SVC = viper.GetString("GRPC_HOST_PRODUCT_SVC")
	GRPC_HOST_CUSTOMER_SVC = viper.GetString("GRPC_HOST_CUSTOMER_SVC")
}
