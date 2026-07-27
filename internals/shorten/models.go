package shorten

import (
	"time"
)

type Nodepath string

const (
	TinyPath  Nodepath = "/tiny"
	blockSize int64    = 10000

	// cache value

	DefaultExpirationTTL = 5 * 365 * 24 * time.Hour
)

type ShortenRequest struct {
	LongURL   string  `json:"long_url"    validate:"required,url"`
	ExpiresAt *string `json:"expires_at"  validate:"omitempty,datetime=2006-01-02T15:04:05Z07:00"`
}
