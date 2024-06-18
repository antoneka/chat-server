package service

import (
	"context"

	"github.com/antoneka/chat-server/internal/model"
)

// ChatService defines the interface for chat-related business logic.
type ChatService interface {
	CreateChat(ctx context.Context, userIDs []int64) (int64, error)
	DeleteChat(ctx context.Context, chatID int64) error
	SendMessage(ctx context.Context, message *model.Message) error
	// ConnectToChat(*ConnectToChatRequest, ChatV1_ConnectToChatServer) error
}
