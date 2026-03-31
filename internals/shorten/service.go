package shorten

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/Jidetireni/tiny/config"
	redis_cache "github.com/Jidetireni/tiny/pkg/Redis"
	"github.com/Jidetireni/tiny/pkg/zookeeper"
	"github.com/google/uuid"
)

var _ Repository = (*ShortenRepository)(nil)
var _ ZookeeperService = (*zookeeper.Zookeeper)(nil)
var _ RedisCacheService = (*redis_cache.RedisCache)(nil)

type ZookeeperService interface {
	GetNextRange(path string, blockSize int64) (int64, int64, error)
}

type Repository interface {
	Create(ctx context.Context, s ShortenedURL) error
}

type RedisCacheService interface {
	Get(ctx context.Context, key string, dest any) (any, error)
	Set(ctx context.Context, key string, value any, ttl time.Duration) error
}

type Service struct {
	Zookeeper  ZookeeperService
	Repo       Repository
	RedisCache RedisCacheService
	Config     *config.Config
	Mu         sync.Mutex
	CurrentID  int64
	RangeEnd   int64
}

func New(
	config *config.Config,
	zookeeper ZookeeperService,
	repo Repository,
	redisCache RedisCacheService,
) *Service {
	return &Service{
		Zookeeper:  zookeeper,
		Repo:       repo,
		RedisCache: redisCache,
	}
}

func (s *Service) Shorten(ctx context.Context, longURL string) (string, error) {
	id, err := s.getNextID()
	if err != nil {
		return "", err
	}

	uniqueCode := base62Encode(uint64(id))

	if err := s.Repo.Create(ctx, ShortenedURL{
		ID:         uuid.New(),
		UniqueCode: uniqueCode,
		LongURL:    longURL,
		CreatedAt:  time.Now(),
	}); err != nil {
		return "", err
	}

	_ = s.RedisCache.Set(ctx, RedisUniqueCodeKey(uniqueCode), longURL, UniqueCodeExpirationTTL)
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
