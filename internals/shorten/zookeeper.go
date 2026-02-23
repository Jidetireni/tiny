package shorten

import (
	"errors"
	"strconv"

	"github.com/go-zookeeper/zk"
)

// getNextRange atomically claims a range of IDs from Zookeeper using
// optimistic concurrency (CAS via zNode versioning).
func getNextRange(conn *zk.Conn, path string, blockSize int64) (int64, int64, error) {
	for {
		data, _, err := conn.Get(path)
		if err != nil {
			if errors.Is(err, zk.ErrNoNode) {
				_, createErr := conn.Create(path, []byte("0"), 0, zk.WorldACL(zk.PermAll))
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

		_, stat, _ := conn.Exists(path)
		_, err = conn.Set(path, []byte(strconv.FormatInt(end, 10)), stat.Version)
		if err != nil {
			if errors.Is(err, zk.ErrBadVersion) {
				continue
			}
			return 0, 0, err
		}

		return start, end, nil
	}
}
