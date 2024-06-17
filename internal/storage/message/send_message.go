package message

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	servicemodel "github.com/antoneka/chat-server/internal/model"
	"github.com/antoneka/chat-server/internal/storage/message/converter"
)

func (s *store) SendMessage(
	ctx context.Context,
	message *servicemodel.Message,
) error {
	storeMessage := converter.ServiceMessageToStoreMessage(message)

	builder := sq.Insert(tableMessages).
		Columns(senderIDColumn, chatIDColumn, messageColumn, sentAtColumn).
		Values(storeMessage.FromUserID, storeMessage.ChatID, storeMessage.Text, storeMessage.SendTime).
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
