package chat

import (
	"context"
	"fmt"
	"github.com/antoneka/chat-server/internal/model"
)

func (s *serv) SendMessage(
	ctx context.Context,
	message *model.Message,
) error {
	const op = "service.chat.SendMessage"

	isUserInChat, err := s.chatMemberStorage.IsUserInChat(ctx, message.ChatID, message.FromUserID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if !isUserInChat {
		return fmt.Errorf("%s: the user %d is not in the chat %d", op, message.FromUserID, message.ChatID)
	}

	err = s.messageStorage.SendMessage(ctx, message)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
