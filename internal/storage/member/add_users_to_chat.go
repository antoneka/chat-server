package member

import (
	"context"
	sq "github.com/Masterminds/squirrel"
)

func (s *store) AddUsersToChat(
	ctx context.Context,
	chatID int64,
	userIDs []int64,
) error {
	builder := sq.Insert(tableChatMembers).
		Columns(chatIDColumn, userIDColumn).
		PlaceholderFormat(sq.Dollar)

	for _, userID := range userIDs {
		builder = builder.Values(chatID, userID)
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
