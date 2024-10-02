package useCaseClients

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/uxsnap/auto_repair/backend/internal/body"
	"github.com/uxsnap/auto_repair/backend/internal/entity"
	"github.com/uxsnap/auto_repair/backend/internal/usecase"
)

type ClientsService struct {
	repo ClientsRepository
}

func NewClientsService(repo ClientsRepository) *ClientsService {
	return &ClientsService{
		repo: repo,
	}
}

func (cs *ClientsService) GetAll(ctx context.Context) ([]entity.Client, error) {
	log.Println("clients: calling GetAll usecase")

	clients := []entity.Client{}

	err := cs.repo.GetAll(ctx, entity.Client{}, clients)

	return clients, err
}

func (cs *ClientsService) Create(ctx context.Context, clientData body.CreateClientBody) (uuid.UUID, error) {
	log.Println("clients: calling Create usecase")

	if err := uuid.Validate(clientData.EmployeeId.String()); err != nil || clientData.EmployeeId == uuid.Nil {
		return uuid.Nil, fmt.Errorf("employeeId must be UUID")
	}

	if len(clientData.Name) < 3 {
		return uuid.Nil, fmt.Errorf("name must be the length of 3 min")
	}

	if !usecase.IsValidPhoneNumber(clientData.Phone) {
		return uuid.Nil, fmt.Errorf("phone is not valid")
	}

	if !usecase.IsValidPass(clientData.Passport) {
		return uuid.Nil, fmt.Errorf("passport is not valid")
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
