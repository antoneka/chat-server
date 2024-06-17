package member

import (
	"context"
	"database/sql"
	"errors"
	sq "github.com/Masterminds/squirrel"
)

func (s *store) IsUserInChat(
	ctx context.Context,
	chatID int64,
	userID int64,
) (bool, error) {
	builder := sq.Select(userIDColumn).
		From(tableChatMembers).
		Where(sq.Eq{chatIDColumn: chatID}, sq.Eq{userIDColumn: userID}).
		Limit(1).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return false, err
	}

	var memberID int64
	err = s.db.QueryRow(ctx, query, args...).Scan(&memberID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
