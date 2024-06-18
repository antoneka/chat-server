package chat

import (
	"github.com/antoneka/chat-server/internal/storage/postgres"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	tableChats = "chats"

	idColumn = "id"
)

var _ postgres.ChatStorage = (*store)(nil)

type store struct {
	db *pgxpool.Pool
}

func NewStorage(db *pgxpool.Pool) postgres.ChatStorage {
	return &store{db: db}
}
