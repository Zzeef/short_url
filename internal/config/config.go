package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port string
	DB   struct {
		Host     string
		User     string
		Password string
		Name     string
	}
}

func LoadConfig() Config {
	viper.SetConfigFile("internal/config/config.env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	var cfg Config
	cfg.Port = viper.GetString("PORT")
	cfg.DB.Host = viper.GetString("DB_HOST")
	cfg.DB.User = viper.GetString("DB_USER")
	cfg.DB.Password = viper.GetString("DB_PASSWORD")
	cfg.DB.Name = viper.GetString("DB_NAME")

	return cfg
}
