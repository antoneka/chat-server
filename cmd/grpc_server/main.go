package main

import (
	"context"
	"fmt"
	"github.com/antoneka/chat-server/internal/converter"
	"github.com/antoneka/chat-server/internal/service"
	servicechat "github.com/antoneka/chat-server/internal/service/chat"
	storagechat "github.com/antoneka/chat-server/internal/storage/chat"
	storagemember "github.com/antoneka/chat-server/internal/storage/member"
	storagemessage "github.com/antoneka/chat-server/internal/storage/message"
	storageuser "github.com/antoneka/chat-server/internal/storage/user"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/antoneka/chat-server/internal/config"
	desc "github.com/antoneka/chat-server/pkg/chat_v1"
	"github.com/jackc/pgx/v4/pgxpool"
)

type server struct {
	desc.UnimplementedChatV1Server
	serv service.ChatService
}

// CreateChat creates a new chat.
func (s *server) CreateChat(
	ctx context.Context,
	req *desc.CreateChatRequest,
) (*desc.CreateChatResponse, error) {
	chatID, err := s.serv.CreateChat(ctx, req.GetUserIds())
	if err != nil {
		return nil, err
	}

	fmt.Printf("Created chat with id: %d\n", chatID)

	return &desc.CreateChatResponse{
		ChatId: chatID,
	}, nil
}

// DeleteChat deletes the chat from the system.
func (s *server) DeleteChat(
	ctx context.Context,
	req *desc.DeleteChatRequest,
) (*emptypb.Empty, error) {
	err := s.serv.DeleteChat(ctx, req.GetChatId())
	if err != nil {
		return nil, err
	}

	fmt.Printf("Deleted chat with id %d\n", req.GetChatId())

	return &emptypb.Empty{}, nil
}

// SendMessage sends a message to the server.
func (s *server) SendMessage(
	ctx context.Context,
	req *desc.SendMessageRequest,
) (*emptypb.Empty, error) {
	err := s.serv.SendMessage(ctx, converter.SendMessageToServiceMessage(req))
	if err != nil {
		return nil, err
	}

	fmt.Printf("Send message %s to chat %d by %d\n", req.Message.GetText(), req.GetChatId(), req.Message.GetFromUserId())

	return &emptypb.Empty{}, nil
}

func main() {
	ctx := context.Background()
	cfg := config.MustLoad()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.GRPC.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pool, err := pgxpool.Connect(ctx, cfg.PG.DSN)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &server{
		serv: servicechat.NewService(
			storagechat.NewStorage(pool),
			storageuser.NewStorage(pool),
			storagemember.NewStorage(pool),
			storagemessage.NewStorage(pool),
		),
	})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve the server: %v", err)
	}
}
