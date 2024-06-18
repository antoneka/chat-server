package chat

import (
	"context"
	"fmt"
)

func (s *serv) DeleteChat(
	ctx context.Context,
	chatID int64,
) error {
	const op = "service.chat.DeleteChat"

	err := s.chatStorage.DeleteChat(ctx, chatID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
