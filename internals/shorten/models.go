package shorten

import "time"

type Nodepath string

const (
	TinyPath  Nodepath = "/tiny"
	blockSize int64    = 1000
)

type ShortenedURL struct {
	ID        string    `json:"id"`
	ShortCode string    `json:"short_code"`
	LongURL   string    `json:"long_url"`
	ExpiresAt time.Time `json:"expires_at"`
}

type ShortenRequest struct {
	LongURL   string `json:"long_url"    validate:"required,url"`
	ExpiresAt string `json:"expires_at"  validate:"omitempty,datetime=2006-01-02T15:04:05Z07:00"`
}
