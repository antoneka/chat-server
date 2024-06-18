package message

import (
	"github.com/antoneka/chat-server/internal/client/db"
	"github.com/antoneka/chat-server/internal/storage/postgres"
)

const (
	tableMessages = "messages"

	idColumn       = "id"
	messageColumn  = "message"
	chatIDColumn   = "chat_id"
	senderIDColumn = "sender_id"
	sentAtColumn   = "sent_at"
)

var _ postgres.MessageStorage = (*store)(nil)

type store struct {
	db db.Client
}

func NewStorage(db db.Client) postgres.MessageStorage {
	return &store{db: db}
}
