package shorten

import (
	"context"

	"github.com/Jidetireni/tiny/pkg/database/cassandra"
)

type ShortenRepository struct {
	db *cassandra.Cassandra
}

func NewShortenRepository(db *cassandra.Cassandra) *ShortenRepository {
	return &ShortenRepository{
		db: db,
	}
}

func (sr *ShortenRepository) Create(ctx context.Context, s ShortenedURL) error {
	query := `INSERT INTO shortened_urls (id, short_code, long_url, expires_at) VALUES (?, ?, ?, ?)`
	return sr.db.ExecuteQuery(ctx, query, s.ID, s.ShortCode, s.LongURL, s.ExpiresAt)
}
