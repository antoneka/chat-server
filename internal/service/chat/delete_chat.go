package chat

import (
	"context"
)

func (s *serv) DeleteChat(
	ctx context.Context,
	chatID int64,
) error {
	err := s.chatStorage.DeleteChat(ctx, chatID)
	if err != nil {
		return err
	}

	return nil
}
