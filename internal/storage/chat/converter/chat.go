package converter

import (
	servicemodel "github.com/antoneka/chat-server/internal/model"
	"github.com/antoneka/chat-server/internal/storage/chat/model"
)

func ServiceChatInfoToStorageChatInfo(chatInfo *servicemodel.ChatInfo) *model.ChatInfo {
	return &model.ChatInfo{
		CreatorUserID: chatInfo.CreatorUserID,
		ChatID:        chatInfo.ChatID,
		ChatName:      chatInfo.ChatName,
	}
}
