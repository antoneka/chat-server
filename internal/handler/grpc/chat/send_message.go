package chat

import (
	"context"
	"fmt"
	"github.com/antoneka/chat-server/internal/handler/grpc/chat/converter"
	desc "github.com/antoneka/chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Implementation) SendMessage(
	ctx context.Context,
	req *desc.SendMessageRequest,
) (*emptypb.Empty, error) {
	const op = "handler.grpc.chat.SendMessage"

	err := s.chatService.SendMessage(ctx, converter.SendMessageToServiceMessage(req))
	if err != nil {
		return &emptypb.Empty{}, fmt.Errorf("%s: %w", op, err)
	}

	return &emptypb.Empty{}, nil
}
