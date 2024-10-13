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

type ClientsRepository interface {
	GetAll(ctx context.Context, params body.ClientBodyParams) ([]entity.ClientWithData, error)
	Create(ctx context.Context, client entity.Client) (uuid.UUID, error)
	Delete(ctx context.Context, clientID string) (uuid.UUID, error)
	Update(ctx context.Context, id uuid.UUID, client entity.Client) error
}

type ClientsService struct {
	repo ClientsRepository
}

func NewClientsService(repo ClientsRepository) *ClientsService {
	return &ClientsService{
		repo: repo,
	}
}

func (cs *ClientsService) GetAll(ctx context.Context, params body.ClientBodyParams) ([]entity.ClientWithData, error) {
	log.Println("clients: calling GetAll usecase")

	return cs.repo.GetAll(ctx, params)
}

func (cs *ClientsService) Create(ctx context.Context, clientData body.CreateClientBody) (uuid.UUID, error) {
	log.Println("clients: calling Create usecase")

	if !validators.IsValidGuid(clientData.EmployeeId) {
		return uuid.Nil, fmt.Errorf("employeeId должен быть UUID")
	}

	if !validators.IsValidLen(clientData.Name, 3) {
		return uuid.Nil, fmt.Errorf("длина имени должна быть больше 3 символов")
	}

	if !validators.IsValidPhoneNumber(clientData.Phone) {
		return uuid.Nil, fmt.Errorf("неверный формат номера")
	}

	if !validators.IsValidPass(clientData.Passport) {
		return uuid.Nil, fmt.Errorf("неверный формат паспорта")
	}

	return cs.repo.Create(ctx, clientData.ToEntity())
}

func (cs *ClientsService) Update(ctx context.Context, id uuid.UUID, clientData body.CreateClientBody) error {
	log.Println("clients: calling Update usecase")

	return cs.repo.Update(ctx, id, clientData.ToEntity())
}

func (cs *ClientsService) Delete(ctx context.Context, clientID uuid.UUID) (uuid.UUID, error) {
	log.Println("clients: calling Delete usecase")

	if clientID == uuid.Nil {
		return uuid.Nil, fmt.Errorf("id must be provided")
	}

	return cs.repo.Delete(ctx, clientID.String())
}
