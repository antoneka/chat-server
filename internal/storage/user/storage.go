package user

import (
	"github.com/antoneka/chat-server/internal/storage"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	tableUsers = "users"

	idColumn = "id"
)

var _ storage.UserStorage = (*store)(nil)

type store struct {
	db *pgxpool.Pool
}

func NewStorage(db *pgxpool.Pool) storage.UserStorage {
	return &store{db: db}
}
