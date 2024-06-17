package chat

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	servicemodel "github.com/antoneka/chat-server/internal/model"
	"github.com/antoneka/chat-server/internal/storage/chat/converter"
)

func (s *store) CreateChat(
	ctx context.Context,
	chatInfo *servicemodel.ChatInfo,
) (int64, error) {
	storeChatInfo := converter.ServiceChatInfoToStorageChatInfo(chatInfo)

	builder := sq.Insert(tableChats).
		Columns(titleColumn, creatorIDColumn).
		Values(storeChatInfo.ChatName, storeChatInfo.CreatorUserID).
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
