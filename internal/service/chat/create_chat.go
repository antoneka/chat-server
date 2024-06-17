package chat

import (
	"context"
	"github.com/antoneka/chat-server/internal/model"
)

func (s *serv) CreateChat(
	ctx context.Context,
	chatInfo *model.ChatInfo,
) (int64, error) {
	chatID, err := s.chatStorage.CreateChat(ctx, chatInfo)
	if err != nil {
		return 0, err
	}

	return chatID, nil
}
