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
	const op = "handler.grpc.chat.DeleteChat"

	err := s.chatService.DeleteChat(ctx, req.GetChatId())
	if err != nil {
		return &emptypb.Empty{}, fmt.Errorf("%s: %w", op, err)
	}

	return &emptypb.Empty{}, nil
}
