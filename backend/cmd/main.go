package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/uxsnap/auto_repair/backend/internal/app"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed loading env %v", err)
		return
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	app := app.New()

	log.Println("App is started")

	app.Run(ctx)
}
