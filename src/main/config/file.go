package config

import (
	"log"

	"github.com/spf13/viper"

)

// LoadFileConfig ...
func LoadFileConfig() {
	viper.SetConfigName("env")
	viper.SetConfigType("json")
	viper.AddConfigPath("./src/main/config/")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panic("Fatal error in load config file:", err)
	}
}

// GetConfig ...
func GetConfig() map[string]interface{} {
	config := viper.GetStringMap("config")
	return config
}
