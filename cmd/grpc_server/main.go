package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/antoneka/chat-server/pkg/chat_v1"
)

const grpcPort = 50052

type server struct {
	desc.UnimplementedChatV1Server
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
		Id: 1337,
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
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve the server: %v", err)
	}
}
