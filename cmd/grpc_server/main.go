package main

import (
	"context"
	"fmt"
	"github.com/antoneka/chat-server/internal/storage"
	"github.com/antoneka/chat-server/internal/storage/chat"
	"github.com/antoneka/chat-server/internal/storage/member"
	"github.com/antoneka/chat-server/internal/storage/message"
	"github.com/antoneka/chat-server/internal/storage/user"
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
	chatRepo       storage.ChatStorage
	userRepo       storage.UserStorage
	chatMemberRepo storage.ChatMemberStorage
	messageRepo    storage.MessageStorage
}

// Create creates a new chat.
func (s *server) Create(
	ctx context.Context,
	req *desc.CreateRequest,
) (*desc.CreateResponse, error) {
	err := s.userRepo.CreateUsers(ctx, []int64{req.GetCreatorUserId()})

	chatId, _ := s.chatRepo.CreateChat(ctx, req)

	if err != nil {
		return nil, err
	}

	fmt.Printf("Created chat with id: %d\n", chatId)

	return &desc.CreateResponse{
		ChatId: chatId,
	}, nil
}

// Delete deletes the chat from the system.
func (s *server) Delete(
	ctx context.Context,
	req *desc.DeleteRequest,
) (*emptypb.Empty, error) {
	err := s.chatRepo.DeleteChat(ctx, req)
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
	err := s.userRepo.CreateUsers(ctx, []int64{req.GetUserId()})

	err = s.messageRepo.SendMessage(ctx, req)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Send message %s to chat %d by %d\n", req.GetText(), req.GetChatId(), req.GetUserId())

	return &emptypb.Empty{}, nil
}

func (s *server) AddUsers(
	ctx context.Context,
	req *desc.AddUsersRequest,
) (*emptypb.Empty, error) {
	err := s.userRepo.CreateUsers(ctx, req.UserIds)

	err = s.chatMemberRepo.AddUsers(ctx, req)

	if err != nil {
		return nil, err
	}

	fmt.Printf("Added users %v to the chat %d\n", req.GetUserIds(), req.GetChatId())

	return &emptypb.Empty{}, nil
}

func (s *server) DeleteUsers(
	ctx context.Context,
	req *desc.DeleteUsersRequest,
) (*emptypb.Empty, error) {
	err := s.chatMemberRepo.DeleteUsers(ctx, req)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Deleted users %v from chat %d\n", req.GetUserIds(), req.GetChatId())

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
		chatRepo:       chat.NewStorage(pool),
		chatMemberRepo: member.NewStorage(pool),
		messageRepo:    message.NewStorage(pool),
		userRepo:       user.NewStorage(pool),
	})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve the server: %v", err)
	}
}
