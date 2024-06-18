package chat

import (
	"github.com/antoneka/chat-server/internal/client/db"
	"github.com/antoneka/chat-server/internal/storage/postgres"
)

const (
	tableChats = "chats"

	idColumn = "id"
)

var _ postgres.ChatStorage = (*store)(nil)

type store struct {
	db db.Client
}

func NewStorage(db db.Client) postgres.ChatStorage {
	return &store{db: db}
}
