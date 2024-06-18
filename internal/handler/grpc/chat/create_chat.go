package chat

import (
	"context"
	"fmt"
	desc "github.com/antoneka/chat-server/pkg/chat_v1"
)

func (s *Implementation) CreateChat(
	ctx context.Context,
	req *desc.CreateChatRequest,
) (*desc.CreateChatResponse, error) {
	const op = "handler.grpc.chat.CreateChat"

	chatID, err := s.chatService.CreateChat(ctx, req.GetUserIds())
	if err != nil {
		return &desc.CreateChatResponse{}, fmt.Errorf("%s: %w", op, err)
	}

	return &desc.CreateChatResponse{
		ChatId: chatID,
	}, nil
}
