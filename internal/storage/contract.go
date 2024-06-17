package storage

import (
	"context"
	servicemodel "github.com/antoneka/chat-server/internal/model"
	desc "github.com/antoneka/chat-server/pkg/chat_v1"
)

type UserStorage interface {
	CreateUsers(ctx context.Context, userIDs []int64) error
}

type ChatStorage interface {
	CreateChat(ctx context.Context, chatInfo *servicemodel.ChatInfo) (int64, error)
	DeleteChat(ctx context.Context, chatInfo *servicemodel.ChatInfo) error
}

type ChatMemberStorage interface {
	AddUsers(ctx context.Context, req *desc.AddUsersRequest) error
	KickUsers(ctx context.Context, req *desc.KickUsersRequest) error
	IsUserInChat(ctx context.Context, chatID int64, userID int64) (bool, error)
}

type MessageStorage interface {
	SendMessage(ctx context.Context, message *servicemodel.Message) error
}
