package database

import (
	"github.com/Jidetireni/tiny/config"
	"github.com/Jidetireni/tiny/pkg/database/cassandra"
)

type Database struct {
	Cassandra *cassandra.Cassandra
}

func New(config *config.Config) (*Database, error) {
	cassandra, err := cassandra.New(&config.CassandraConfig)
	if err != nil {
		return nil, err
	}

	return &Database{
		Cassandra: cassandra,
	}, nil
}
