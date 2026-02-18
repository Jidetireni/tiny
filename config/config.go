package config

import (
	"log"
	"os"
)

type ServerConfig struct {
	Host string
	Port string
}

type DatabaseConfig struct {
	URL  string
	Type string
}

type ZooKeeperConfig struct {
	Host string
	Port string
}

type RedisConfig struct {
	URL string
}

func validateENV() {
	environmentVariables := []string{}

	for _, env := range environmentVariables {
		if os.Getenv(env) == "" {
			log.Fatalf("Environment variable %s is not set", env)
		}
	}
}

type Config struct {
	ServerConfig
	DatabaseConfig
	ZooKeeperConfig
	RedisConfig
}

func New() *Config {

	validateENV()

	return &Config{
		ServerConfig:    ServerConfig{},
		DatabaseConfig:  DatabaseConfig{},
		ZooKeeperConfig: ZooKeeperConfig{},
		RedisConfig:     RedisConfig{},
	}
}
