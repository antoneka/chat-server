package chat

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/antoneka/chat-server/internal/handler/grpc/chat/converter"
	desc "github.com/antoneka/chat-server/pkg/chat_v1"
)

// SendMessage handles the gRPC request to send a message in a chat.
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
