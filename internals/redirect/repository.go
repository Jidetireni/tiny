package redirect

import (
	"context"

	"github.com/Jidetireni/tiny/internals"
	"github.com/Jidetireni/tiny/pkg/database/cassandra"
)

type RedirectRepository struct {
	db *cassandra.Cassandra
}

func NewRedirectRepository(db *cassandra.Cassandra) *RedirectRepository {
	return &RedirectRepository{
		db: db,
	}
}

func (rr *RedirectRepository) GetByShortCode(ctx context.Context, shortCode string) (*internals.ShortenedURL, error) {
	query := `SELECT * FROM redirects WHERE unique_code = ?`
	var sh internals.ShortenedURL
	err := rr.db.ExecuteQuery(ctx, query, &sh, shortCode)
	if err != nil {
		return nil, err
	}
	return &sh, nil
}
