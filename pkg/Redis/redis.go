package redis_cache

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/Jidetireni/tiny/config"
	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	Client *redis.Client
}

func New(config *config.Config) (*RedisCache, error) {
	opts, err := redis.ParseURL(config.RedisConfig.URL)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opts)

	return &RedisCache{
		Client: client,
	}, nil
}

// Get fetches a JSON string from Redis and deserializes it into the 'dest' pointer.
func (r *RedisCache) Get(ctx context.Context, key string, dest any) (any, error) {
	valueStr, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			// add log
			return nil, nil
		}
		return nil, err
	}

	err = json.Unmarshal([]byte(valueStr), dest)
	if err != nil {
		return nil, err
	}

	return dest, nil
}

// Set stores a JSON string in Redis with the given key and TTL.
func (r *RedisCache) Set(ctx context.Context, key string, value any, TTL time.Duration) error {
	valueStr, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = r.Client.Set(ctx, key, valueStr, TTL).Err()
	if err != nil {
		return err
	}

	return nil
}
