package converter

import (
	"github.com/antoneka/chat-server/internal/model"
	desc "github.com/antoneka/chat-server/pkg/chat_v1"
)

// SendMessageToServiceMessage converts a gRPC send message request to a service layer message model.
func SendMessageToServiceMessage(sendMessageRequest *desc.SendMessageRequest) *model.Message {
	grpcMessage := sendMessageRequest.GetMessage()

	return &model.Message{
		ChatID:     sendMessageRequest.GetChatId(),
		FromUserID: grpcMessage.GetFromUserId(),
		Text:       grpcMessage.GetText(),
		SendTime:   grpcMessage.GetSentAt().AsTime(),
	}
}
