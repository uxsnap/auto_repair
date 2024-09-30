package app

import (
	"context"
	"log"

	"github.com/uxsnap/auto_repair/backend/internal/config"
	"github.com/uxsnap/auto_repair/backend/internal/db"
	repoClients "github.com/uxsnap/auto_repair/backend/internal/repo/client"
	useCaseClients "github.com/uxsnap/auto_repair/backend/internal/usecase/client"
)

func (a *App) GetConfigDB() *config.ConfigDB {
	if a.configDB == nil {
		a.configDB = config.NewConfigDB()
	}
	return a.configDB
}

func (a *App) GetConfigHTTP() *config.ConfigHTTP {
	if a.configHTTP == nil {
		a.configHTTP = config.NewConfigHttp()
	}
	return a.configHTTP
}

func (a *App) DbClient(ctx context.Context) *db.Client {
	if a.db == nil {
		client, err := db.New(ctx, a.configDB.DSN())

		if err != nil {
			log.Fatalf("failed to connect to postgres: %v", err)
		}

		a.db = client
	}

	return a.db
}

func (a *App) ClientsRepository(ctx context.Context) *repoClients.ClientsRepository {
	if a.repoClients == nil {
		a.repoClients = repoClients.NewClientsRepo(a.DbClient(ctx))
	}
	return a.repoClients
}

func (a *App) ClientsService(ctx context.Context) *useCaseClients.ClientsService {
	if a.ucClients == nil {
		a.ucClients = useCaseClients.NewClientsService(
			a.ClientsRepository(ctx),
		)
	}

	return a.ucClients
}
