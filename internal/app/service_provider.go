package app

import (
	"context"
	"log"

	"github.com/antoneka/chat-server/internal/config"
	"github.com/antoneka/chat-server/internal/handler/grpc/chat"
	"github.com/antoneka/chat-server/internal/service"
	chatService "github.com/antoneka/chat-server/internal/service/chat"
	"github.com/antoneka/chat-server/internal/storage/postgres"
	chatStorage "github.com/antoneka/chat-server/internal/storage/postgres/chat"
	memberStorage "github.com/antoneka/chat-server/internal/storage/postgres/member"
	messageStorage "github.com/antoneka/chat-server/internal/storage/postgres/message"
	userStorage "github.com/antoneka/chat-server/internal/storage/postgres/user"
	"github.com/antoneka/chat-server/pkg/client/db"
	"github.com/antoneka/chat-server/pkg/client/db/pg"
	"github.com/antoneka/chat-server/pkg/client/db/transaction"
	"github.com/antoneka/chat-server/pkg/closer"
)

// serviceProvider is a DI container that manages service dependencies.
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

// newServiceProvider creates a new instance of serviceProvider.
func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// Config returns the configuration settings for the application.
func (s *serviceProvider) Config() *config.Config {
	if s.config == nil {
		s.config = config.MustLoad()
	}

	return s.config
}

// DBClient returns the database client for interacting with the database.
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

// TxManager returns the transaction manager for handling database transactions.
func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

// ChatStorage returns the chat storage interface for accessing chat-related data.
func (s *serviceProvider) ChatStorage(ctx context.Context) postgres.ChatStorage {
	if s.chatStore == nil {
		s.chatStore = chatStorage.NewStorage(s.DBClient(ctx))
	}

	return s.chatStore
}

// MessageStorage returns the message storage interface for accessing message-related data.
func (s *serviceProvider) MessageStorage(ctx context.Context) postgres.MessageStorage {
	if s.messageStore == nil {
		s.messageStore = messageStorage.NewStorage(s.DBClient(ctx))
	}

	return s.messageStore
}

// MemberStorage returns the member storage interface for accessing chat member-related data.
func (s *serviceProvider) MemberStorage(ctx context.Context) postgres.ChatMemberStorage {
	if s.memberStore == nil {
		s.memberStore = memberStorage.NewStorage(s.DBClient(ctx))
	}

	return s.memberStore
}

// UserStorage returns the user storage interface for accessing user-related data.
func (s *serviceProvider) UserStorage(ctx context.Context) postgres.UserStorage {
	if s.userStore == nil {
		s.userStore = userStorage.NewStorage(s.DBClient(ctx))
	}

	return s.userStore
}

// ChatService returns the chat service for managing chat operations.
func (s *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if s.chatService == nil {
		s.chatService = chatService.NewService(
			s.ChatStorage(ctx),
			s.UserStorage(ctx),
			s.MemberStorage(ctx),
			s.MessageStorage(ctx),
			s.TxManager(ctx),
		)
	}

	return s.chatService
}

// ChatHandler returns the chat handler for handling chat-related API requests.
func (s *serviceProvider) ChatHandler(ctx context.Context) *chat.Implementation {
	if s.chatHandler == nil {
		s.chatHandler = chat.NewImplementation(s.ChatService(ctx))
	}

	return s.chatHandler
}
