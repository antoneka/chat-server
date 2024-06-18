package chat

import (
	"context"
	"fmt"
	"github.com/antoneka/chat-server/internal/converter"
	desc "github.com/antoneka/chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Implementation) SendMessage(
	ctx context.Context,
	req *desc.SendMessageRequest,
) (*emptypb.Empty, error) {
	err := s.chatService.SendMessage(ctx, converter.SendMessageToServiceMessage(req))
	if err != nil {
		return &emptypb.Empty{}, err
	}

	fmt.Printf("Send message %s to chat %d by %d\n", req.Message.GetText(), req.GetChatId(), req.Message.GetFromUserId())

	return &emptypb.Empty{}, nil
}
