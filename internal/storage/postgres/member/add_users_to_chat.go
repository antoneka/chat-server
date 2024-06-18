package member

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/antoneka/chat-server/internal/client/db"
)

func (s *store) AddUsersToChat(
	ctx context.Context,
	chatID int64,
	userIDs []int64,
) error {
	const op = "storage.postgres.member.AddUsersToChat"

	builder := sq.Insert(tableChatMembers).
		Columns(chatIDColumn, userIDColumn).
		PlaceholderFormat(sq.Dollar)

	for _, userID := range userIDs {
		builder = builder.Values(chatID, userID)
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
