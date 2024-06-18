package member

import (
	"github.com/antoneka/chat-server/internal/client/db"
	"github.com/antoneka/chat-server/internal/storage/postgres"
)

const (
	tableChatMembers = "chat_members"

	chatIDColumn = "chat_id"
	userIDColumn = "user_id"
)

var _ postgres.ChatMemberStorage = (*store)(nil)

type store struct {
	db db.Client
}

func NewStorage(db db.Client) postgres.ChatMemberStorage {
	return &store{db: db}
}
