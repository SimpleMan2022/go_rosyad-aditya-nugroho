package utils

import (
	"github.com/spf13/viper"
)

func GetConfig(key string) string {
	viper.SetConfigFile(".env")
	viper.AddConfigPath(".")

	//if err := viper.ReadInConfig(); err != nil {
	//	log.Fatalf("error when reading configuration file: %s\n", err)
	//}
	return viper.GetString(key)
}
