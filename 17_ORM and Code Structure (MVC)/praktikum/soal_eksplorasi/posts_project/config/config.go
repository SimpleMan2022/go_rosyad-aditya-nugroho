package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_URL      string
	DB_NAME     string
	DB_PORT     string
}

var ENV *Config

func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigName(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	if err := viper.Unmarshal(&ENV); err != nil {
		log.Fatal(err)
	}
}
