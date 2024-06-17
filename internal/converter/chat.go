package converter

import (
	"github.com/antoneka/chat-server/internal/model"
	desc "github.com/antoneka/chat-server/pkg/chat_v1"
)

func SendMessageToServiceMessage(sendMessageRequest *desc.SendMessageRequest) *model.Message {
	grpcMessage := sendMessageRequest.GetMessage()

	return &model.Message{
		ChatID:     sendMessageRequest.GetChatId(),
		FromUserID: grpcMessage.GetFromUserId(),
		Text:       grpcMessage.GetText(),
		SendTime:   grpcMessage.GetSentAt().AsTime(),
	}
}

func CreateChatToServiceChatInfo(createChatRequest *desc.CreateChatRequest) *model.ChatInfo {
	return &model.ChatInfo{
		CreatorUserID: createChatRequest.GetCreatorUserId(),
		ChatName:      createChatRequest.GetChatName(),
	}
}

func DeleteChatToServiceChatInfo(deleteChatRequest *desc.DeleteChatRequest) *model.ChatInfo {
	return &model.ChatInfo{
		CreatorUserID: deleteChatRequest.GetCreatorId(),
		ChatID:        deleteChatRequest.GetChatId(),
	}
}

func AddUsersToServiceParams(addUsersRequest *desc.AddUsersRequest) *model.AddUsersParam {
	chatInfo := model.ChatInfo{
		CreatorUserID: addUsersRequest.GetCreatorId(),
		ChatID:        addUsersRequest.GetChatId(),
	}

	return &model.AddUsersParam{
		ChatInfo: chatInfo,
		UserIDs:  addUsersRequest.GetUserIds(),
	}
}

func KickUsersToServiceParams(kickUsersRequest *desc.KickUsersRequest) *model.KickUsersParam {
	chatInfo := model.ChatInfo{
		CreatorUserID: kickUsersRequest.GetCreatorId(),
		ChatID:        kickUsersRequest.GetChatId(),
	}

	return &model.KickUsersParam{
		ChatInfo: chatInfo,
		UserIDs:  kickUsersRequest.GetUserIds(),
	}
}
