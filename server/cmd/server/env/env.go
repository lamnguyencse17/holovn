package env

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

func ReadEnv(key string) string {
	if os.Getenv("PRODUCTION") == "TRUE" || os.Getenv("TESTING") == "CI" {
		value := os.Getenv(key)
		if value == "" {
			log.Fatalln(key, " NOT FOUND IN ENV")
		}
		return value
	}
	if os.Getenv("TESTING") == "LOCAL" {
		viper.SetConfigFile("../.env")
	} else {
		viper.SetConfigFile("./cmd/server/.env")
	}
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Error while reading config file %s\n", err)
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	return value
}
