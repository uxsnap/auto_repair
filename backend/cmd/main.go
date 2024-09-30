package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/uxsnap/auto_repair/backend/internal/app"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	app := app.New()

	log.Println("App is started")

	app.Run(ctx)
}
