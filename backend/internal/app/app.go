package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/uxsnap/auto_repair/backend/internal/config"
	handler "github.com/uxsnap/auto_repair/backend/internal/delivery/http"
)

type App struct {
	handler    *handler.Handler
	configHTTP *config.ConfigHTTP
}

func New() *App {
	h := handler.New()

	return &App{
		handler:    h,
		configHTTP: config.NewConfigHttp(),
	}
}

func (a *App) Run(ctx context.Context) {
	ch := make(chan error, 1)

	configHttp := config.NewConfigHttp()

	server := http.Server{
		Addr:    configHttp.Host + ":" + configHttp.Port,
		Handler: a.handler.Router,
	}

	go func() {
		fmt.Printf("Server is listening on port %v \n", a.configHTTP.Port)

		err := server.ListenAndServe()

		if err != nil {
			ch <- err
		}

		close(ch)
	}()

	go func() {
		select {
		case err := <-ch:
			log.Println("Unexpected error:")
			log.Fatal(err)
		case <-ctx.Done():
			log.Printf("Server is shutting down")

			timeoutCtx, cancel := context.WithTimeout(ctx, time.Second*10)
			defer cancel()

			if err := server.Shutdown(timeoutCtx); err != nil {
				log.Printf("http server shutdown error %v", err)
			}
		}
	}()
}
