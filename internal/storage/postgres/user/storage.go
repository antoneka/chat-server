package user

import (
	"github.com/antoneka/chat-server/internal/client/db"
	"github.com/antoneka/chat-server/internal/storage/postgres"
)

const (
	tableUsers = "users"

	idColumn = "id"
)

var _ postgres.UserStorage = (*store)(nil)

type store struct {
	db db.Client
}

func NewStorage(db db.Client) postgres.UserStorage {
	return &store{db: db}
}
