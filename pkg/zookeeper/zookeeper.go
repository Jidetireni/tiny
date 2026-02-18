package zookeeper

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/Jidetireni/tiny/config"
	"github.com/go-zookeeper/zk"
)

type Zookeeper struct {
	conn *zk.Conn
}

func New(config *config.Config) (*Zookeeper, error) {
	server := fmt.Sprintf("%s:%s", config.ZooKeeperConfig.Host, config.ZooKeeperConfig.Port)

	conn, _, err := zk.Connect([]string{server}, time.Second)
	if err != nil {
		return nil, fmt.Errorf("zookeeper: failed to connect to %s: %w", server, err)
	}

	return &Zookeeper{
		conn: conn,
	}, nil
}

func (z *Zookeeper) GetNextRange(path string, blockSize int64) (int64, int64, error) {
	for {
		data, _, err := z.conn.Get(path)
		if err != nil {
			if errors.Is(err, zk.ErrNoNode) {
				_, createErr := z.conn.Create(path, []byte("0"), 0, zk.WorldACL(zk.PermAll))
				if createErr != nil && !errors.Is(createErr, zk.ErrNodeExists) {
					return 0, 0, err
				}
				continue
			}
			return 0, 0, err
		}

		start, err := strconv.ParseInt(string(data), 10, 64)
		if err != nil {
			return 0, 0, err
		}

		end := start + blockSize

		_, stat, _ := z.conn.Exists(path)
		_, err = z.conn.Set(path, []byte(strconv.FormatInt(end, 10)), stat.Version)
		if err != nil {
			if errors.Is(err, zk.ErrBadVersion) {
				continue
			}
			return 0, 0, err
		}

		return start, end, nil
	}
}
