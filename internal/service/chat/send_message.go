package chat

import (
	"context"
	"github.com/antoneka/chat-server/internal/model"
)

func (s *serv) SendMessage(
	ctx context.Context,
	message *model.Message,
) error {
	err := s.messageStorage.SendMessage(ctx, message)
	if err != nil {
		return err
	}

	return nil
}
