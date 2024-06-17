package member

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	servicemodel "github.com/antoneka/chat-server/internal/model"
)

func (s *store) AddUsers(
	ctx context.Context,
	addUsersParam *servicemodel.AddUsersParam,
) error {
	builder := sq.Insert(tableChatMembers).
		Columns(chatIDColumn, userIDColumn).
		PlaceholderFormat(sq.Dollar)

	userIDs := req.GetUserIds()
	chatID := req.GetChatId()

	for _, userID := range userIDs {
		builder = builder.Values(chatID, userID)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
