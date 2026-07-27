package shorten

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/Jidetireni/tiny/config"
	"github.com/Jidetireni/tiny/internals"
	"github.com/Jidetireni/tiny/pkg/zookeeper"
	"github.com/google/uuid"
)

var _ Repository = (*ShortenRepository)(nil)
var _ ZookeeperService = (*zookeeper.Zookeeper)(nil)

type ZookeeperService interface {
	GetNextRange(path string, blockSize int64) (int64, int64, error)
}

type Repository interface {
	Create(ctx context.Context, s internals.ShortenedURL) error
}

type Service struct {
	Zookeeper ZookeeperService
	Repo      Repository
	Config    *config.Config
	Mu        sync.Mutex
	CurrentID int64
	RangeEnd  int64
}

func New(
	config *config.Config,
	zookeeper ZookeeperService,
	repo Repository,
) *Service {
	return &Service{
		Zookeeper: zookeeper,
		Repo:      repo,
	}
}

func (s *Service) Shorten(ctx context.Context, req *ShortenRequest) (string, error) {
	id, err := s.getNextID()
	if err != nil {
		return "", err
	}

	uniqueCode := base62Encode(uint64(id))

	expiresAt := time.Now().Add(DefaultExpirationTTL)
	if req.ExpiresAt != nil {
		exp, err := time.Parse(time.RFC3339, *req.ExpiresAt)
		if err == nil {
			expiresAt = exp
		}
	}

	if err := s.Repo.Create(ctx, internals.ShortenedURL{
		ID:         uuid.New(),
		UniqueCode: uniqueCode,
		LongURL:    req.LongURL,
		ExpiresAt:  expiresAt,
		CreatedAt:  time.Now(),
	}); err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s", s.Config.BaseURL, uniqueCode), nil
}

func (s *Service) getNextID() (int64, error) {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	if s.CurrentID >= s.RangeEnd {
		start, end, err := s.Zookeeper.GetNextRange(string(TinyPath), blockSize)
		if err != nil {
			return 0, err
		}
		s.CurrentID = start
		s.RangeEnd = end
	}

	id := s.CurrentID
	s.CurrentID++
	return id, nil
}
