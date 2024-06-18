package chat

import (
	"context"
	"fmt"
)

// CreateChat handles the creation of a new chat.
func (s *serv) CreateChat(
	ctx context.Context,
	userIDs []int64,
) (int64, error) {
	const op = "service.chat.CreateChat"

	var chatID int64

	err := s.txManager.ReadCommitted(ctx, func(context.Context) error {
		var errTx error
		chatID, errTx = s.chatStorage.CreateChat(ctx)
		if errTx != nil {
			return errTx
		}

		errTx = s.userStorage.CreateUsers(ctx, userIDs)
		if errTx != nil {
			return errTx
		}

		errTx = s.chatMemberStorage.AddUsersToChat(ctx, chatID, userIDs)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return chatID, nil
}
