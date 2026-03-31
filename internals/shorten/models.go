package shorten

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Nodepath string

type RedisKeyPrefix string

const (
	TinyPath  Nodepath = "/tiny"
	blockSize int64    = 10000

	// cache value
	RedisUniqueCodeKeyPrefix RedisKeyPrefix = "unique_code:%s" //uniqueCode
	UniqueCodeExpirationTTL  time.Duration  = 30 * time.Minute
)

type ShortenedURL struct {
	ID         uuid.UUID `json:"id"`
	UniqueCode string    `json:"unique_code"`
	LongURL    string    `json:"long_url"`
	ExpiresAt  time.Time `json:"expires_at"`
	CreatedAt  time.Time `json:"created_at"`
}

type ShortenRequest struct {
	LongURL   string `json:"long_url"    validate:"required,url"`
	ExpiresAt string `json:"expires_at"  validate:"omitempty,datetime=2006-01-02T15:04:05Z07:00"`
}

func RedisUniqueCodeKey(uniqueCode string) string {
	return fmt.Sprintf(string(RedisUniqueCodeKeyPrefix), uniqueCode)
}
