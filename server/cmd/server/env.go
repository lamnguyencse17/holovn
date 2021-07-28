package main

import (
	"github.com/spf13/viper"
	"log"
)

func ReadEnv(key string) string {
	viper.SetConfigFile("./cmd/server/.env")
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