package chat

import (
	"github.com/antoneka/platform-common/pkg/db"

	"github.com/antoneka/chat-server/internal/service"
	"github.com/antoneka/chat-server/internal/storage/postgres"
)

var _ service.ChatService = (*serv)(nil)

// serv is the implementation of the ChatService interface.
type serv struct {
	chatStorage       postgres.ChatStorage
	userStorage       postgres.UserStorage
	chatMemberStorage postgres.ChatMemberStorage
	messageStorage    postgres.MessageStorage

	txManager db.TxManager
}

// NewService creates a new instance of serv and returns it as a ChatService.
func NewService(
	chatStorage postgres.ChatStorage,
	userStorage postgres.UserStorage,
	chatMemberStorage postgres.ChatMemberStorage,
	messageStorage postgres.MessageStorage,
	txManager db.TxManager,
) service.ChatService {
	return &serv{
		chatStorage:       chatStorage,
		userStorage:       userStorage,
		chatMemberStorage: chatMemberStorage,
		messageStorage:    messageStorage,
		txManager:         txManager,
	}
}
