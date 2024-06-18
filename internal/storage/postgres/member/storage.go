package member

import (
	"github.com/antoneka/chat-server/internal/storage/postgres"
	"github.com/antoneka/chat-server/pkg/client/db"
)

const (
	tableChatMembers = "chat_members"

	chatIDColumn = "chat_id"
	userIDColumn = "user_id"
)

var _ postgres.ChatMemberStorage = (*store)(nil)

// store represents the implementation of the ChatMemberStorage interface.
type store struct {
	db db.Client
}

// NewStorage creates a new instance of ChatMemberStorage.
func NewStorage(db db.Client) postgres.ChatMemberStorage {
	return &store{db: db}
}
