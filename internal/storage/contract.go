package storage

import (
	"context"
	servicemodel "github.com/antoneka/chat-server/internal/model"
)

type UserStorage interface {
	CreateUsers(ctx context.Context, userIDs []int64) error
}

type ChatStorage interface {
	CreateChat(ctx context.Context) (int64, error)
	DeleteChat(ctx context.Context, chatID int64) error
}

type ChatMemberStorage interface {
	AddUsersToChat(ctx context.Context, chatID int64, userIDs []int64) error
	IsUserInChat(ctx context.Context, chatID int64, userID int64) (bool, error)
}

type MessageStorage interface {
	SendMessage(ctx context.Context, message *servicemodel.Message) error
}
