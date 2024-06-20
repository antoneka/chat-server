package message

import (
	"github.com/antoneka/platform-common/pkg/db"

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

// store represents the implementation of the MessageStorage interface.
type store struct {
	db db.Client
}

// NewStorage creates a new instance of MessageStorage.
func NewStorage(db db.Client) postgres.MessageStorage {
	return &store{db: db}
}
