package utils

import (
	"log"
	"otp_service/structs"

	"github.com/spf13/viper"
)

var ApplicationConfig structs.Config

func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	// Find and read the config file
	err := viper.ReadInConfig()

	if err != nil {
		log.Printf("Error while reading config file %s", err)
	}

	err = viper.Unmarshal(&ApplicationConfig)

	if err != nil {
		log.Printf("Error while reading config file %s", err)
	}
}
