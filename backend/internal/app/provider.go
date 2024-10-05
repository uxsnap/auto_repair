package app

import (
	"context"
	"log"

	"github.com/uxsnap/auto_repair/backend/internal/config"
	"github.com/uxsnap/auto_repair/backend/internal/db"
	"github.com/uxsnap/auto_repair/backend/internal/repo"
	"github.com/uxsnap/auto_repair/backend/internal/usecase"
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

func (a *App) ClientsRepository(ctx context.Context) *repo.ClientsRepository {
	if a.repoClients == nil {
		a.repoClients = repo.NewClientsRepo(a.DbClient(ctx))
	}
	return a.repoClients
}

func (a *App) ClientsService(ctx context.Context) *usecase.ClientsService {
	if a.ucClients == nil {
		a.ucClients = usecase.NewClientsService(
			a.ClientsRepository(ctx),
		)
	}

	return a.ucClients
}

func (a *App) EmployeesRepository(ctx context.Context) *repo.EmployeesRepository {
	if a.repoEmployees == nil {
		a.repoEmployees = repo.NewEmployeesRepo(a.DbClient(ctx))
	}
	return a.repoEmployees
}

func (a *App) EmployeesService(ctx context.Context) *usecase.EmployeesService {
	if a.ucEmployees == nil {
		a.ucEmployees = usecase.NewEmployeesService(
			a.EmployeesRepository(ctx),
		)
	}

	return a.ucEmployees
}

func (a *App) ContractsRepository(ctx context.Context) *repo.ContractsRepository {
	if a.repoContracts == nil {
		a.repoContracts = repo.NewContractsRepo(a.DbClient(ctx))
	}
	return a.repoContracts
}

func (a *App) ContractsService(ctx context.Context) *usecase.ContractsService {
	if a.ucContracts == nil {
		a.ucContracts = usecase.NewContractsService(
			a.ContractsRepository(ctx),
		)
	}

	return a.ucContracts
}

func (a *App) ReceiptsRepository(ctx context.Context) *repo.ReceiptsRepository {
	if a.repoReceipts == nil {
		a.repoReceipts = repo.NewReceiptsRepo(a.DbClient(ctx))
	}
	return a.repoReceipts
}

func (a *App) ReceiptsService(ctx context.Context) *usecase.ReceiptsService {
	if a.ucReceipts == nil {
		a.ucReceipts = usecase.NewReceiptsService(
			a.ReceiptsRepository(ctx),
		)
	}

	return a.ucReceipts
}

func (a *App) VehiclesRepository(ctx context.Context) *repo.VehiclesRepository {
	if a.repoVehicles == nil {
		a.repoVehicles = repo.NewVehiclesRepo(a.DbClient(ctx))
	}
	return a.repoVehicles
}

func (a *App) VehiclesService(ctx context.Context) *usecase.VehiclesService {
	if a.ucVehicles == nil {
		a.ucVehicles = usecase.NewVehiclesService(
			a.VehiclesRepository(ctx),
		)
	}

	return a.ucVehicles
}

func (a *App) ApplicationsRepository(ctx context.Context) *repo.ApplicationsRepository {
	if a.repoApplications == nil {
		a.repoApplications = repo.NewApplicationsRepo(a.DbClient(ctx))
	}
	return a.repoApplications
}

func (a *App) ApplicationsService(ctx context.Context) *usecase.ApplicationsService {
	if a.ucApplications == nil {
		a.ucApplications = usecase.NewApplicationsService(
			a.ApplicationsRepository(ctx),
		)
	}

	return a.ucApplications
}
