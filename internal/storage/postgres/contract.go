package postgres

import (
	"context"

	servicemodel "github.com/antoneka/chat-server/internal/model"
)

// UserStorage defines the interface for user-related database operations.
type UserStorage interface {
	CreateUsers(ctx context.Context, userIDs []int64) error
}

// ChatStorage defines the interface for chat-related database operations.
type ChatStorage interface {
	CreateChat(ctx context.Context) (int64, error)
	DeleteChat(ctx context.Context, chatID int64) error
}

// ChatMemberStorage defines the interface for chat member-related database operations.
type ChatMemberStorage interface {
	AddUsersToChat(ctx context.Context, chatID int64, userIDs []int64) error
	IsUserInChat(ctx context.Context, chatID int64, userID int64) (bool, error)
}

// MessageStorage defines the interface for message-related database operations.
type MessageStorage interface {
	SendMessage(ctx context.Context, message *servicemodel.Message) error
}
