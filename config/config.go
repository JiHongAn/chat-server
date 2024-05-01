package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Jwt      JwtConfig
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	DatabaseUrl string
}

type JwtConfig struct {
	AccessTokenSecret string
}

var AppConfig Config

func LoadConfig() {
	if _, err := toml.DecodeFile("config.toml", &AppConfig); err != nil {
		log.Fatal("Error loading config.toml:", err)
	}
}
