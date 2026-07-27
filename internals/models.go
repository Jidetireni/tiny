package internals

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

const (
	RedisUniqueCodeKeyPrefix = "unique_code:%s" //uniqueCode
	UniqueCodeExpirationTTL  = 30 * time.Minute
)

type ShortenedURL struct {
	ID         uuid.UUID `json:"id"`
	UniqueCode string    `json:"unique_code"`
	LongURL    string    `json:"long_url"`
	ExpiresAt  time.Time `json:"expires_at"`
	CreatedAt  time.Time `json:"created_at"`
}

func RedisUniqueCodeKey(uniqueCode string) string {
	return fmt.Sprintf(string(RedisUniqueCodeKeyPrefix), uniqueCode)
}
