package user

import (
	"github.com/antoneka/chat-server/internal/storage/postgres"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	tableUsers = "users"

	idColumn = "id"
)

var _ postgres.UserStorage = (*store)(nil)

type store struct {
	db *pgxpool.Pool
}

func NewStorage(db *pgxpool.Pool) postgres.UserStorage {
	return &store{db: db}
}
