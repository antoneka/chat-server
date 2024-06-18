package member

import (
	"context"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/antoneka/chat-server/internal/client/db"
	"github.com/jackc/pgx/v4"
)

func (s *store) IsUserInChat(
	ctx context.Context,
	chatID int64,
	userID int64,
) (bool, error) {
	const op = "storage.postgres.member.IsUserInChat"

	builder := sq.Select(userIDColumn).
		From(tableChatMembers).
		Where(sq.Eq{chatIDColumn: chatID, userIDColumn: userID}).
		Limit(1).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	var memberID int64
	err = s.db.DB().ScanOneContext(ctx, &memberID, q, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}

		return false, fmt.Errorf("%s: %w", op, err)
	}

	return true, nil
}
