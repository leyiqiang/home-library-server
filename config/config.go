package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	Server   Server
	Database Database
	Auth     Auth
}

type Database struct {
	URI      string
	Name     string
	Username string
	Password string
}

type Server struct {
	Port string
}

type Auth struct {
	Secret string
}

func (c *Config) Read() {

	viper.SetConfigType("yml")
	if os.Getenv("ENV") == "prod" {
		viper.SetConfigName("config-prod")
	} else {
		viper.SetConfigName("config")
	}
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config, %s", err)
	}
	err := viper.Unmarshal(&c)
	if err != nil {
		log.Fatalf("Error decoding config, %v", err)
	}
}
