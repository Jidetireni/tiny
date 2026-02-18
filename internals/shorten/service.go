package shorten

import (
	"sync"

	"github.com/Jidetireni/tiny/pkg/zookeeper"
)

var _ ZookeeperStore = (*zookeeper.Zookeeper)(nil)

type ZookeeperStore interface {
	GetNextRange(path string, blockSize int64) (int64, int64, error)
}

type Service struct {
	zk        ZookeeperStore
	mu        sync.Mutex
	currentID int64
	rangeEnd  int64
}

func New(zk ZookeeperStore) *Service {
	return &Service{zk: zk}
}

func (s *Service) Shorten(longURL string) (string, error) {
	id, err := s.nextID()
	if err != nil {
		return "", err
	}

	shortCode := encode(id)
	_ = shortCode // TODO: store longURL → shortCode mapping in DB
	return shortCode, nil
}

func (s *Service) nextID() (int64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.currentID >= s.rangeEnd {
		start, end, err := s.zk.GetNextRange(string(TinyPath), blockSize)
		if err != nil {
			return 0, err
		}
		s.currentID = start
		s.rangeEnd = end
	}

	id := s.currentID
	s.currentID++
	return id, nil
}
