package message

import (
	"github.com/antoneka/chat-server/internal/storage"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	tableMessages = "messages"

	idColumn       = "id"
	messageColumn  = "message"
	chatIDColumn   = "chat_id"
	senderIDColumn = "sender_id"
	sentAtColumn   = "sent_at"
)

var _ storage.MessageStorage = (*store)(nil)

type store struct {
	db *pgxpool.Pool
}

func NewStorage(db *pgxpool.Pool) storage.MessageStorage {
	return &store{db: db}
}
