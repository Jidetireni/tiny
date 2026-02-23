package shorten

import (
	"sync"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
	"github.com/go-zookeeper/zk"
)

type Service struct {
	zkConn    *zk.Conn
	session   *gocql.Session
	mu        sync.Mutex
	currentID int64
	rangeEnd  int64
}

func New(zkConn *zk.Conn, session *gocql.Session) *Service {
	return &Service{
		zkConn:  zkConn,
		session: session,
	}
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
		start, end, err := getNextRange(s.zkConn, string(TinyPath), blockSize)
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
