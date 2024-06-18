package message

import (
	"github.com/antoneka/chat-server/internal/storage/postgres"
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

var _ postgres.MessageStorage = (*store)(nil)

type store struct {
	db *pgxpool.Pool
}

func NewStorage(db *pgxpool.Pool) postgres.MessageStorage {
	return &store{db: db}
}
