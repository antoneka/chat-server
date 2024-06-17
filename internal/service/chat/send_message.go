package chat

import (
	"context"
	"errors"
	"github.com/antoneka/chat-server/internal/model"
)

func (s *serv) SendMessage(
	ctx context.Context,
	message *model.Message,
) error {
	isUserInChat, err := s.chatMemberStorage.IsUserInChat(ctx, message.ChatID, message.FromUserID)
	if err != nil {
		return err
	}

	if !isUserInChat {
		return errors.New("user is not in chat")
	}

	err = s.messageStorage.SendMessage(ctx, message)
	if err != nil {
		return err
	}

	return nil
}
