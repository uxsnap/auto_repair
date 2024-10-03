package app

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/uxsnap/auto_repair/backend/internal/config"
	"github.com/uxsnap/auto_repair/backend/internal/db"
	handler "github.com/uxsnap/auto_repair/backend/internal/delivery/http"
	"github.com/uxsnap/auto_repair/backend/internal/repo"
	"github.com/uxsnap/auto_repair/backend/internal/usecase"
)

type App struct {
	Handler    *handler.Handler
	configHTTP *config.ConfigHTTP
	configDB   *config.ConfigDB

	db            *db.Client
	repoClients   *repo.ClientsRepository
	repoEmployees *repo.EmployeesRepository
	repoContracts *repo.ContractsRepository

	ucClients   *usecase.ClientsService
	ucEmployees *usecase.EmployeesService
	ucContracts *usecase.ContractsService
}

func New() *App {
	return &App{
		configHTTP: config.NewConfigHttp(),
		configDB:   config.NewConfigDB(),
	}
}

func (a *App) Run(ctx context.Context) {
	a.Handler = handler.New(
		a.ClientsService(ctx),
		a.EmployeesService(ctx),
		a.ContractsService(ctx),
	)

	ch := make(chan error, 1)

	server := http.Server{
		Addr:    a.configHTTP.Addr(),
		Handler: a.Handler.Router,
	}

	go func() {
		log.Printf("Server is listening on %v \n", a.configHTTP.Addr())

		err := server.ListenAndServe()

		if err != nil {
			log.Printf("Error while starting server %v \n", err)
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

	<-ctx.Done()
}
