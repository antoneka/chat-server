package chat

import (
	"context"
)

func (s *serv) CreateChat(
	ctx context.Context,
	userIDs []int64,
) (int64, error) {
	chatID, err := s.chatStorage.CreateChat(ctx)
	if err != nil {
		return 0, err
	}

	err = s.userStorage.CreateUsers(ctx, userIDs)
	if err != nil {
		return 0, err
	}

	err = s.chatMemberStorage.AddUsersToChat(ctx, chatID, userIDs)
	if err != nil {
		return 0, err
	}

	return chatID, nil
}
