package chat

import (
	"github.com/antoneka/chat-server/internal/service"
	"github.com/antoneka/chat-server/internal/storage/postgres"
)

var _ service.ChatService = (*serv)(nil)

type serv struct {
	chatStorage       postgres.ChatStorage
	userStorage       postgres.UserStorage
	chatMemberStorage postgres.ChatMemberStorage
	messageStorage    postgres.MessageStorage
}

func NewService(
	chatStorage postgres.ChatStorage,
	userStorage postgres.UserStorage,
	chatMemberStorage postgres.ChatMemberStorage,
	messageStorage postgres.MessageStorage,
) service.ChatService {
	return &serv{
		chatStorage:       chatStorage,
		userStorage:       userStorage,
		chatMemberStorage: chatMemberStorage,
		messageStorage:    messageStorage,
	}
}
