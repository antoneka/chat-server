package service

import (
	"context"
	"github.com/antoneka/chat-server/internal/model"
)

type ChatService interface {
	CreateChat(ctx context.Context, chatInfo *model.ChatInfo) (int64, error)
	DeleteChat(ctx context.Context, chatInfo *model.ChatInfo) error
	SendMessage(ctx context.Context, message *model.Message) error
	AddUsers(ctx context.Context, addUsersParam *model.AddUsersParam) error
	KickUsers(ctx context.Context, kickUsersParam *model.KickUsersParam) error
	// ConnectToChat(*ConnectToChatRequest, ChatV1_ConnectToChatServer) error
}
