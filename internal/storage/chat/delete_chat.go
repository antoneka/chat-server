package chat

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	servicemodel "github.com/antoneka/chat-server/internal/model"
	"github.com/antoneka/chat-server/internal/storage/chat/converter"
	"github.com/pkg/errors"
)

func (s *store) DeleteChat(
	ctx context.Context,
	chatInfo *servicemodel.ChatInfo,
) error {
	storeChatInfo := converter.ServiceChatInfoToStorageChatInfo(chatInfo)

	creatorID, err := s.findChatCreator(ctx, storeChatInfo.CreatorUserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("incorrect ID of the chat creator")
		}
		return err
	}

	builder := sq.Delete(tableChats).
		Where(sq.Eq{idColumn: creatorID}).
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

func (s *store) findChatCreator(
	ctx context.Context,
	chatID int64,
) (int64, error) {
	builder := sq.Select(creatorIDColumn).
		From(tableChats).
		Where(sq.Eq{idColumn: chatID}).
		Limit(1).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	var creatorID int64
	err = s.db.QueryRow(ctx, query, args...).Scan(&creatorID)
	if err != nil {
		return 0, err
	}

	return creatorID, nil
}
