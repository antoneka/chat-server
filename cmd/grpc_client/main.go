package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	desc "github.com/antoneka/chat-server/pkg/chat_v1"
	"github.com/fatih/color"
)

const address = "localhost:50051"

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to create a connection: %v", err)
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Fatalf("failed to close the connection: %v", err)
		}
	}()

	client := desc.NewChatV1Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	resp, err := client.Create(ctx, &desc.CreateRequest{
		Usernames: []string{"antoneka1", "antoneka2"},
	})
	if err != nil {
		log.Fatalf("failed to create a chat: %v", err)
	}

	log.Println(color.RedString("ID of the created chat: %v", resp.GetId()))
}
