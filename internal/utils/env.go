package utils

import (
	"log"

	"github.com/spf13/viper"
)

// return the value of the key
func ViperEnvVariable(key string) string {

	viper.SetConfigFile("../.env")

	// Find and read the config file
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	value, ok := viper.Get(key).(string)

	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return value
}
