package user

import (
	"github.com/antoneka/chat-server/internal/storage/postgres"
	"github.com/antoneka/chat-server/pkg/client/db"
)

const (
	tableUsers = "users"

	idColumn = "id"
)

var _ postgres.UserStorage = (*store)(nil)

// store represents the implementation of the UserStorage interface.
type store struct {
	db db.Client
}

// NewStorage creates a new instance of UserStorage.
func NewStorage(db db.Client) postgres.UserStorage {
	return &store{db: db}
}
