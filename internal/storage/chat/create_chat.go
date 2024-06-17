package chat

import (
	"context"
	sq "github.com/Masterminds/squirrel"
)

func (s *store) CreateChat(
	ctx context.Context,
) (int64, error) {
	builder := sq.Insert(tableChats).
		Columns(idColumn).
		Values(sq.Expr("DEFAULT")).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	var id int64
	err = s.db.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
