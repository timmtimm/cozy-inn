package util

import (
	"log"

	"github.com/spf13/viper"
)

func GetFirebaseEnv(key string) string {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".firebase/service-account.json")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error when reading config: %s", err)
	}

	return viper.GetString(key)
}

func GetConfig(key string) string {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error when reading config: %s", err)
	}

	return viper.GetString(key)
}
