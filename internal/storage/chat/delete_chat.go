package chat

import (
	"context"
	sq "github.com/Masterminds/squirrel"
)

func (s *store) DeleteChat(
	ctx context.Context,
	chatID int64,
) error {
	builder := sq.Delete(tableChats).
		Where(sq.Eq{idColumn: chatID}).
		PlaceholderFormat(sq.Dollar)

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
