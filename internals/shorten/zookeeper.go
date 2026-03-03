package shorten

import (
	"errors"
	"strconv"

	"github.com/go-zookeeper/zk"
)

// IDGenerator provides unique ID ranges for short-code generation.

// ZookeeperIDGen implements IDGenerator using Zookeeper's
// optimistic concurrency (CAS via zNode versioning).
type ZookeeperRepo struct {
	conn *zk.Conn
}

func NewZookeeper(conn *zk.Conn) *ZookeeperRepo {
	return &ZookeeperRepo{conn: conn}
}

func (z *ZookeeperRepo) GetNextRange(path string, blockSize int64) (int64, int64, error) {
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
