package member

import (
	"github.com/antoneka/chat-server/internal/storage/postgres"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	tableChatMembers = "chat_members"

	chatIDColumn = "chat_id"
	userIDColumn = "user_id"
)

var _ postgres.ChatMemberStorage = (*store)(nil)

type store struct {
	db *pgxpool.Pool
}

func NewStorage(db *pgxpool.Pool) postgres.ChatMemberStorage {
	return &store{db: db}
}
