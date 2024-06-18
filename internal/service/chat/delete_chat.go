package chat

import (
	"context"
	"fmt"
)

// DeleteChat handles the deletion of an existing chat.
func (s *serv) DeleteChat(
	ctx context.Context,
	chatID int64,
) error {
	const op = "service.chat.DeleteChat"

	err := s.txManager.ReadCommitted(ctx, func(context.Context) error {
		errTx := s.chatStorage.DeleteChat(ctx, chatID)
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
