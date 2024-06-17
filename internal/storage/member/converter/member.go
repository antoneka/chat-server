package converter

import (
	servicemodel "github.com/antoneka/chat-server/internal/model"
	"github.com/antoneka/chat-server/internal/storage/member/model"
)

func ServiceAddUsersParamToStorageAddUsersParam(addUsersParam servicemodel.AddUsersParam) *model.AddUsersParam {
	return &model.AddUsersParam{
		ChatID:    addUsersParam.ChatInfo.ChatID,
		CreatorID: addUsersParam.ChatInfo.CreatorUserID,
		UserIDs:   addUsersParam.UserIDs,
	}
}
