package user

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"github.com/antoneka/platform-common/pkg/db"
)

// CreateUsers creates new users in the database.
func (s *store) CreateUsers(
	ctx context.Context,
	userIDs []int64,
) error {
	const op = "storage.postgres.user.CreateUsers"

	builder := sq.Insert(tableUsers).
		Columns(idColumn).
		PlaceholderFormat(sq.Dollar).
		Suffix("ON CONFLICT (id) DO NOTHING")

	for _, userID := range userIDs {
		builder = builder.Values(userID)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	_, err = s.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
