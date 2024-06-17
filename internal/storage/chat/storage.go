package chat

import (
	"github.com/antoneka/chat-server/internal/storage"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	tableChats = "chats"

	idColumn = "id"
)

var _ storage.ChatStorage = (*store)(nil)

type store struct {
	db *pgxpool.Pool
}

func NewStorage(db *pgxpool.Pool) storage.ChatStorage {
	return &store{db: db}
}
