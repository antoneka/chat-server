package chat

import (
	"github.com/antoneka/platform-common/pkg/db"

	"github.com/antoneka/chat-server/internal/storage/postgres"
)

const (
	tableChats = "chats"

	idColumn = "id"
)

var _ postgres.ChatStorage = (*store)(nil)

// store represents the implementation of the ChatStorage interface.
type store struct {
	db db.Client
}

// NewStorage creates a new instance of ChatStorage.
func NewStorage(db db.Client) postgres.ChatStorage {
	return &store{db: db}
}
