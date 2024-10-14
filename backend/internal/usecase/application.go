package usecase

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/uxsnap/auto_repair/backend/internal/body"
	"github.com/uxsnap/auto_repair/backend/internal/entity"
	"github.com/uxsnap/auto_repair/backend/internal/validators"
)

type ApplicationsRepository interface {
	GetAll(ctx context.Context, params body.ApplicationBodyParams) ([]entity.ApplicationWithData, error)
	Create(ctx context.Context, client entity.Application) (uuid.UUID, error)
	Delete(ctx context.Context, clientID string) (uuid.UUID, error)
	Update(ctx context.Context, ID uuid.UUID, clientData entity.Application) error
}

type ApplicationsService struct {
	repo ApplicationsRepository
}

func NewApplicationsService(repo ApplicationsRepository) *ApplicationsService {
	return &ApplicationsService{
		repo: repo,
	}
}

func (cs *ApplicationsService) GetAll(ctx context.Context, params body.ApplicationBodyParams) ([]entity.ApplicationWithData, error) {
	log.Println("Applications: calling GetAll usecase")

	return cs.repo.GetAll(ctx, params)
}

func (cs *ApplicationsService) Create(ctx context.Context, clientData body.CreateApplicationBody) (uuid.UUID, error) {
	log.Println("Applications: calling Create usecase")

	if !validators.IsValidGuid(clientData.EmployeeId) {
		return uuid.Nil, fmt.Errorf("employeeId должен быть UUID")
	}

	if !validators.IsValidGuid(clientData.ClientId) {
		return uuid.Nil, fmt.Errorf("clientId должен быть UUID")
	}

	if !validators.IsValidGuid(clientData.ContractId) {
		return uuid.Nil, fmt.Errorf("contractId должен быть UUID")
	}

	if !validators.IsValidLen(clientData.Name, 10) {
		return uuid.Nil, fmt.Errorf("длина имени должна быть больше 10 символов")
	}

	if !validators.IsValidLen(clientData.Status, 1) {
		return uuid.Nil, fmt.Errorf("длина статуса должна быть не равна нулю")
	}

	return cs.repo.Create(ctx, clientData.ToEntity())
}

func (cs *ApplicationsService) Delete(ctx context.Context, clientID uuid.UUID) (uuid.UUID, error) {
	log.Println("Applications: calling Delete usecase")

	if clientID == uuid.Nil {
		return uuid.Nil, fmt.Errorf("id must be provided")
	}

	return cs.repo.Delete(ctx, clientID.String())
}

func (cs *ApplicationsService) Update(ctx context.Context, id uuid.UUID, clientData body.CreateApplicationBody) error {
	log.Println("Applications: calling Update usecase")

	return cs.repo.Update(ctx, id, clientData.ToEntity())
}
