package main

import (
	"context"
	"fmt"
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
	pool *pgxpool.Pool
}

// Create creates a new chat.
func (s *server) Create(
	ctx context.Context,
	req *desc.CreateRequest,
) (*desc.CreateResponse, error) {
	_, cancel := context.WithCancel(ctx)
	defer cancel()

	fmt.Printf("CreateRequest: %+v\n", req)

	return &desc.CreateResponse{
		ChatId: 1337,
	}, nil
}

// Delete deletes the chat from the system.
func (s *server) Delete(
	ctx context.Context,
	req *desc.DeleteRequest,
) (*emptypb.Empty, error) {
	_, cancel := context.WithCancel(ctx)
	defer cancel()

	fmt.Printf("DeleteRequest: %+v\n", req)

	return &emptypb.Empty{}, nil
}

// SendMessage sends a message to the server.
func (s *server) SendMessage(
	ctx context.Context,
	req *desc.SendMessageRequest,
) (*emptypb.Empty, error) {
	_, cancel := context.WithCancel(ctx)
	defer cancel()

	fmt.Printf("SendMessageRequest: %+v\n", req)

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
	desc.RegisterChatV1Server(s, &server{pool: pool})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve the server: %v", err)
	}
}
