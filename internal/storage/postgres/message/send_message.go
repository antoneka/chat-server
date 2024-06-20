package message

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"github.com/antoneka/platform-common/pkg/db"

	servicemodel "github.com/antoneka/chat-server/internal/model"
	"github.com/antoneka/chat-server/internal/storage/postgres/message/converter"
)

// SendMessage sends a message to a chat.
func (s *store) SendMessage(
	ctx context.Context,
	message *servicemodel.Message,
) error {
	const op = "storage.postgres.message.SendMessage"

	storeMessage := converter.ServiceMessageToStoreMessage(message)

	builder := sq.Insert(tableMessages).
		Columns(senderIDColumn, chatIDColumn, messageColumn, sentAtColumn).
		Values(storeMessage.FromUserID, storeMessage.ChatID, storeMessage.Text, storeMessage.SendTime).
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
