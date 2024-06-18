package main

import (
	"context"
	"github.com/antoneka/chat-server/internal/app"
	"log"
)

func main() {
	application, err := app.NewApp(context.Background())
	if err != nil {
		log.Fatalf("failed to init app: %v", err)
	}

	err = application.Run()
	if err != nil {
		log.Fatalf("failed to run app: %v", err)
	}
}
