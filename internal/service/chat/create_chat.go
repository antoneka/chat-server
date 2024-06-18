package chat

import (
	"context"
	"fmt"
)

func (s *serv) CreateChat(
	ctx context.Context,
	userIDs []int64,
) (int64, error) {
	const op = "service.chat.CreateChat"

	chatID, err := s.chatStorage.CreateChat(ctx)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	err = s.userStorage.CreateUsers(ctx, userIDs)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	err = s.chatMemberStorage.AddUsersToChat(ctx, chatID, userIDs)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return chatID, nil
}
