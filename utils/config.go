package utils

import (
	"dcard-2024-backend-intern-assignment/configs"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	NAME        string
	VERSION     string
	PORT        string
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
}

func setDefaultConfigs() {
	configs.SetServerConfigs()
	configs.SetDatabaseConfigs()
}

func LoadConfig() Config {
	setDefaultConfigs()

	viper.SetConfigName("env")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	return config
}
