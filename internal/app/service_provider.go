package app

import (
	"context"
	"github.com/antoneka/chat-server/internal/client/db"
	"github.com/antoneka/chat-server/internal/client/db/pg"
	"github.com/antoneka/chat-server/internal/client/db/transaction"
	"github.com/antoneka/chat-server/internal/closer"
	"github.com/antoneka/chat-server/internal/config"
	"github.com/antoneka/chat-server/internal/handler/grpc/chat"
	"github.com/antoneka/chat-server/internal/service"
	chatService "github.com/antoneka/chat-server/internal/service/chat"
	"github.com/antoneka/chat-server/internal/storage/postgres"
	chatStorage "github.com/antoneka/chat-server/internal/storage/postgres/chat"
	memberStorage "github.com/antoneka/chat-server/internal/storage/postgres/member"
	messageStorage "github.com/antoneka/chat-server/internal/storage/postgres/message"
	userStorage "github.com/antoneka/chat-server/internal/storage/postgres/user"
	"log"
)

type serviceProvider struct {
	config *config.Config

	dbClient  db.Client
	txManager db.TxManager

	chatStore    postgres.ChatStorage
	userStore    postgres.UserStorage
	memberStore  postgres.ChatMemberStorage
	messageStore postgres.MessageStorage

	chatService service.ChatService

	chatHandler *chat.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) Config() *config.Config {
	if s.config == nil {
		s.config = config.MustLoad()
	}

	return s.config
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		client, err := pg.NewDBClient(ctx, s.Config().PG.DSN)
		if err != nil {
			log.Panicf("failed to create db client %v", err)
		}

		err = client.DB().Ping(ctx)
		if err != nil {
			log.Panicf("ping error: %v", err)
		}

		closer.Add(client.Close)

		s.dbClient = client
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) ChatStorage(ctx context.Context) postgres.ChatStorage {
	if s.chatStore == nil {
		s.chatStore = chatStorage.NewStorage(s.DBClient(ctx))
	}

	return s.chatStore
}

func (s *serviceProvider) MessageStorage(ctx context.Context) postgres.MessageStorage {
	if s.messageStore == nil {
		s.messageStore = messageStorage.NewStorage(s.DBClient(ctx))
	}

	return s.messageStore
}

func (s *serviceProvider) MemberStorage(ctx context.Context) postgres.ChatMemberStorage {
	if s.memberStore == nil {
		s.memberStore = memberStorage.NewStorage(s.DBClient(ctx))
	}

	return s.memberStore
}

func (s *serviceProvider) UserStorage(ctx context.Context) postgres.UserStorage {
	if s.userStore == nil {
		s.userStore = userStorage.NewStorage(s.DBClient(ctx))
	}

	return s.userStore
}

func (s *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if s.chatService == nil {
		s.chatService = chatService.NewService(
			s.ChatStorage(ctx),
			s.UserStorage(ctx),
			s.MemberStorage(ctx),
			s.MessageStorage(ctx),
		)
	}

	return s.chatService
}

func (s *serviceProvider) ChatHandler(ctx context.Context) *chat.Implementation {
	if s.chatHandler == nil {
		s.chatHandler = chat.NewImplementation(s.ChatService(ctx))
	}

	return s.chatHandler
}
