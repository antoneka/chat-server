package chat

import (
	"context"
	"fmt"

	"github.com/antoneka/chat-server/internal/model"
)

// SendMessage handles sending a message to a chat.
func (s *serv) SendMessage(
	ctx context.Context,
	message *model.Message,
) error {
	const op = "service.chat.SendMessage"

	err := s.txManager.ReadCommitted(ctx, func(context.Context) error {
		var errTx error
		isUserInChat, errTx := s.chatMemberStorage.IsUserInChat(ctx, message.ChatID, message.FromUserID)
		if errTx != nil {
			return errTx
		}

		if !isUserInChat {
			return fmt.Errorf("the user %d is not in the chat %d", message.FromUserID, message.ChatID)
		}

		errTx = s.messageStorage.SendMessage(ctx, message)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
