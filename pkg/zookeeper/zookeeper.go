package zookeeper

import (
	"fmt"
	"time"

	"github.com/Jidetireni/tiny/config"
	"github.com/go-zookeeper/zk"
)

type Zookeeper struct {
	Conn *zk.Conn
}

func New(config *config.Config) (*Zookeeper, error) {
	server := fmt.Sprintf("%s:%s", config.ZooKeeperConfig.Host, config.ZooKeeperConfig.Port)

	conn, _, err := zk.Connect([]string{server}, time.Second)
	if err != nil {
		return nil, fmt.Errorf("zookeeper: failed to connect to %s: %w", server, err)
	}

	return &Zookeeper{
		Conn: conn,
	}, nil
}
