package config

import (
	"log"
	"os"
)

type ServerConfig struct {
	Host string
	Port string
}

type CassandraConfig struct {
	Host     string
	KeySpace string
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
	CassandraConfig
	ZooKeeperConfig
	RedisConfig
}

func New() *Config {

	validateENV()

	return &Config{
		ServerConfig:    ServerConfig{},
		CassandraConfig: CassandraConfig{},
		ZooKeeperConfig: ZooKeeperConfig{},
		RedisConfig:     RedisConfig{},
	}
}
