package chat

import (
	"github.com/antoneka/chat-server/internal/service"
	"github.com/antoneka/chat-server/internal/storage"
)

var _ service.ChatService = (*serv)(nil)

type serv struct {
	chatStorage       storage.ChatStorage
	userStorage       storage.UserStorage
	chatMemberStorage storage.ChatMemberStorage
	messageStorage    storage.MessageStorage
}

func NewService(
	chatStorage storage.ChatStorage,
	userStorage storage.UserStorage,
	chatMemberStorage storage.ChatMemberStorage,
	messageStorage storage.MessageStorage,
) service.ChatService {
	return &serv{
		chatStorage:       chatStorage,
		userStorage:       userStorage,
		chatMemberStorage: chatMemberStorage,
		messageStorage:    messageStorage,
	}
}
