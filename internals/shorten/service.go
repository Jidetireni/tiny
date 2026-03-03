package shorten

import (
	"context"
	"sync"
)

var _ IDGenerator = (*ZookeeperRepo)(nil)
var _ Repository = (*ShortenRepository)(nil)

type IDGenerator interface {
	GetNextRange(path string, blockSize int64) (int64, int64, error)
}

type Repository interface {
	Create(ctx context.Context, s ShortenedURL) error
}

type Service struct {
	idGen     IDGenerator
	repo      Repository
	mu        sync.Mutex
	currentID int64
	rangeEnd  int64
}

func New(idGen IDGenerator, repo Repository) *Service {
	return &Service{
		idGen: idGen,
		repo:  repo,
	}
}

func (s *Service) Shorten(ctx context.Context, longURL string) (string, error) {
	id, err := s.nextID()
	if err != nil {
		return "", err
	}

	shortCode := encode(id)

	if err := s.repo.Create(ctx, ShortenedURL{
		ID:        shortCode,
		ShortCode: shortCode,
		LongURL:   longURL,
	}); err != nil {
		return "", err
	}

	return shortCode, nil
}

func (s *Service) nextID() (int64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.currentID >= s.rangeEnd {
		start, end, err := s.idGen.GetNextRange(string(TinyPath), blockSize)
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
