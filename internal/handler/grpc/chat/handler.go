package chat

import (
	"github.com/antoneka/chat-server/internal/service"
	desc "github.com/antoneka/chat-server/pkg/chat_v1"
)

// Implementation is the gRPC handler implementation for chat-related operations.
type Implementation struct {
	desc.UnimplementedChatV1Server
	chatService service.ChatService
}

// NewImplementation creates a new instance of the chat handler Implementation.
func NewImplementation(chatService service.ChatService) *Implementation {
	return &Implementation{
		chatService: chatService,
	}
}
