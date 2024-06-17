package user

import (
	"context"
	sq "github.com/Masterminds/squirrel"
)

func (s *store) CreateUsers(
	ctx context.Context,
	userIDs []int64,
) error {
	builder := sq.Insert(tableUsers).
		Columns(idColumn).
		PlaceholderFormat(sq.Dollar).
		Suffix("ON CONFLICT (id) DO NOTHING")

	for _, userID := range userIDs {
		builder = builder.Values(userID)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	_, err = s.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
