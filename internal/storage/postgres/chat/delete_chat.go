package chat

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"github.com/antoneka/chat-server/pkg/client/db"
)

// DeleteChat deletes a chat from the database using its ID.
func (s *store) DeleteChat(
	ctx context.Context,
	chatID int64,
) error {
	const op = "storage.postgres.chat.DeleteChat"

	builder := sq.Delete(tableChats).
		Where(sq.Eq{idColumn: chatID}).
		PlaceholderFormat(sq.Dollar)

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
