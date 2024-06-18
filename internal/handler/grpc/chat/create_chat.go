package chat

import (
	"context"
	desc "github.com/antoneka/chat-server/pkg/chat_v1"
)

func (s *Implementation) CreateChat(
	ctx context.Context,
	req *desc.CreateChatRequest,
) (*desc.CreateChatResponse, error) {
	chatID, err := s.chatService.CreateChat(ctx, req.GetUserIds())
	if err != nil {
		return &desc.CreateChatResponse{}, err
	}

	return &desc.CreateChatResponse{
		ChatId: chatID,
	}, nil
}
