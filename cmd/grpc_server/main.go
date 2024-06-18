package main

import (
	"context"
	"fmt"
	chatHandler "github.com/antoneka/chat-server/internal/handler/grpc/chat"
	chatService "github.com/antoneka/chat-server/internal/service/chat"
	chatStorage "github.com/antoneka/chat-server/internal/storage/postgres/chat"
	memberStorage "github.com/antoneka/chat-server/internal/storage/postgres/member"
	messageStorage "github.com/antoneka/chat-server/internal/storage/postgres/message"
	userStorage "github.com/antoneka/chat-server/internal/storage/postgres/user"
	desc "github.com/antoneka/chat-server/pkg/chat_v1"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/antoneka/chat-server/internal/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

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

	chatStore := chatStorage.NewStorage(pool)
	userStore := userStorage.NewStorage(pool)
	memberStore := memberStorage.NewStorage(pool)
	messageStore := messageStorage.NewStorage(pool)

	chatSrv := chatService.NewService(chatStore, userStore, memberStore, messageStore)

	desc.RegisterChatV1Server(s, chatHandler.NewImplementation(chatSrv))

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve the server: %v", err)
	}
}
