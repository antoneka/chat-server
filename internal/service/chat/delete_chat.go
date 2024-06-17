package chat

import (
	"context"
	"github.com/antoneka/chat-server/internal/model"
)

func (s *serv) DeleteChat(
	ctx context.Context,
	chatInfo *model.ChatInfo,
) error {
	err := s.chatStorage.DeleteChat(ctx, chatInfo)
	if err != nil {
		return err
	}

	return nil
}
