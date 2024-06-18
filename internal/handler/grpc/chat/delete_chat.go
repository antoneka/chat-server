package chat

import (
	"context"
	"fmt"
	desc "github.com/antoneka/chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Implementation) DeleteChat(
	ctx context.Context,
	req *desc.DeleteChatRequest,
) (*emptypb.Empty, error) {
	err := s.chatService.DeleteChat(ctx, req.GetChatId())
	if err != nil {
		return &emptypb.Empty{}, err
	}

	fmt.Printf("Deleted chat with id %d\n", req.GetChatId())

	return &emptypb.Empty{}, nil
}
