package converter

import (
	servicemodel "github.com/antoneka/chat-server/internal/model"
	"github.com/antoneka/chat-server/internal/storage/message/model"
)

func ServiceMessageToStorageMessage(message *servicemodel.Message) *model.Message {
	return &model.Message{
		ChatID:     message.ChatID,
		FromUserID: message.FromUserID,
		Text:       message.Text,
		SendTime:   message.SendTime,
	}
}
