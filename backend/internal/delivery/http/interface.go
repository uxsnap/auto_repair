package handler

import (
	"context"

	"github.com/google/uuid"
	"github.com/uxsnap/auto_repair/backend/internal/body"
	"github.com/uxsnap/auto_repair/backend/internal/entity"
)

type ClientsService interface {
	GetAll(ctx context.Context) ([]entity.Client, error)
	Create(ctx context.Context, clientData body.CreateClientBody) (uuid.UUID, error)
	Update(ctx context.Context, ID uuid.UUID, clientData body.CreateClientBody) error
	Delete(ctx context.Context, ID uuid.UUID) (uuid.UUID, error)
}

type ReceiptsService interface {
	GetAll(ctx context.Context) ([]entity.Receipt, error)
	Create(ctx context.Context, clientData body.CreateReceiptBody) (uuid.UUID, error)
	Update(ctx context.Context, ID uuid.UUID, clientData body.CreateReceiptBody) error
	Delete(ctx context.Context, ID uuid.UUID) (uuid.UUID, error)
}

type ApplicationsService interface {
	GetAll(ctx context.Context) ([]entity.Application, error)
	Create(ctx context.Context, clientData body.CreateApplicationBody) (uuid.UUID, error)
	Delete(ctx context.Context, ID uuid.UUID) (uuid.UUID, error)
}

type VehiclesService interface {
	GetAll(ctx context.Context) ([]entity.Vehicle, error)
	Create(ctx context.Context, clientData body.CreateVehicleBody) (uuid.UUID, error)
	Delete(ctx context.Context, ID uuid.UUID) (uuid.UUID, error)
}

type EmployeesService interface {
	GetAll(ctx context.Context) ([]entity.Employee, error)
	Create(ctx context.Context, EmployeesData body.CreateEmployeeBody) (uuid.UUID, error)
	Update(ctx context.Context, ID uuid.UUID, EmployeesData body.CreateEmployeeBody) error
}

type ContractsService interface {
	GetAll(ctx context.Context) ([]entity.Contract, error)
	Create(ctx context.Context, ContractsData body.CreateContractBody) (uuid.UUID, error)
	Update(ctx context.Context, ID uuid.UUID, ContractsData body.CreateContractBody) error
}
