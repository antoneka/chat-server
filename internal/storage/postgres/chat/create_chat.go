package chat

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/antoneka/chat-server/internal/client/db"
)

func (s *store) CreateChat(
	ctx context.Context,
) (int64, error) {
	const op = "storage.postgres.chat.CreateChat"

	builder := sq.Insert(tableChats).
		Columns(idColumn).
		Values(sq.Expr("DEFAULT")).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	var id int64
	err = s.db.DB().ScanOneContext(ctx, &id, q, args...)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}
