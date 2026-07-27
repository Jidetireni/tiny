package config

import (
	"log"
	"os"
)

type ServerConfig struct {
	Host    string
	Port    string
	BaseURL string
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

type Config struct {
	ServerConfig
	CassandraConfig
	ZooKeeperConfig
	RedisConfig
}

func validateENV() {
	environmentVariables := []string{
		"HOST",
		"PORT",
		"BASE_URL",
		"CASSANDRA_HOST",
		"CASSANDRA_KEYSPACE",
		"ZOOKEEPER_HOST",
		"ZOOKEEPER_PORT",
		"REDIS_URL",
	}

	for _, env := range environmentVariables {
		if os.Getenv(env) == "" {
			log.Fatalf("Environment variable %s is not set", env)
		}
	}
}

func New() *Config {

	validateENV()

	return &Config{
		ServerConfig: ServerConfig{
			Host:    os.Getenv("HOST"),
			Port:    os.Getenv("PORT"),
			BaseURL: os.Getenv("BASE_URL"),
		},
		CassandraConfig: CassandraConfig{
			Host:     os.Getenv("CASSANDRA_HOST"),
			KeySpace: os.Getenv("CASSANDRA_KEYSPACE"),
		},
		ZooKeeperConfig: ZooKeeperConfig{
			Host: os.Getenv("ZOOKEEPER_HOST"),
			Port: os.Getenv("ZOOKEEPER_PORT"),
		},
		RedisConfig: RedisConfig{
			URL: os.Getenv("REDIS_URL"),
		},
	}
}
