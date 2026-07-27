package redirect

import (
	"context"
	"time"

	"github.com/Jidetireni/tiny/internals"
	redis_cache "github.com/Jidetireni/tiny/pkg/Redis"
)

var _ RedisCacheService = (*redis_cache.RedisCache)(nil)
var _ Repository = (*RedirectRepository)(nil)

type Repository interface {
	GetByShortCode(ctx context.Context, shortCode string) (*internals.ShortenedURL, error)
}

type RedisCacheService interface {
	Get(ctx context.Context, key string, dest any) (any, error)
	Set(ctx context.Context, key string, value any, ttl time.Duration) error
}

type Service struct {
	Repo       Repository
	RedisCache RedisCacheService
}

func New(repo Repository, redisCache RedisCacheService) *Service {
	return &Service{
		Repo:       repo,
		RedisCache: redisCache,
	}
}

func (s *Service) Redirect(ctx context.Context, shortCode string) (string, error) {
	key := internals.RedisUniqueCodeKey(shortCode)

	var longURL string
	if _, err := s.RedisCache.Get(ctx, key, &longURL); err == nil && longURL != "" {
		return longURL, nil // cache hit
	}

	// cache miss → hit Cassandra
	sh, err := s.Repo.GetByShortCode(ctx, shortCode)
	if err != nil {
		return "", err
	}

	_ = s.RedisCache.Set(ctx, key, sh.LongURL, internals.UniqueCodeExpirationTTL) // populate for next time
	return sh.LongURL, nil
}
