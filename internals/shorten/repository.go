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
	query := `INSERT INTO shortened_urls (id, unique_code, long_url, expires_at, created_at) VALUES (?, ?, ?, ?, ?)`
	return sr.db.ExecuteQuery(ctx, query, s.ID, s.UniqueCode, s.LongURL, s.ExpiresAt, s.CreatedAt)
}
