package useCaseClients

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
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
	return cs.repo.GetAll(ctx)
}

func (cs *ClientsService) Create(ctx context.Context, clientData entity.CreateClientBody) error {
	log.Println("clients: calling Create usecase")

	if err := uuid.Validate(clientData.EmployeeId.String()); err != nil || clientData.EmployeeId == uuid.Nil {
		return fmt.Errorf("employeeId must be UUID")
	}

	if len(clientData.Name) < 3 {
		return fmt.Errorf("name must be the length of 3 min")
	}

	if !usecase.IsValidPhoneNumber(clientData.Phone) {
		return fmt.Errorf("phone is not valid")
	}

	if !usecase.IsValidPass(clientData.Passport) {
		return fmt.Errorf("passport is not valid")
	}

	return cs.repo.Create(ctx, clientData.ToEntity())
}

func (cs *ClientsService) Update(ctx context.Context, clientData entity.Client) error {
	log.Println("clients: calling Update usecase")
	return nil
}

func (cs *ClientsService) Delete(ctx context.Context, clientID pgtype.UUID) error {
	log.Println("clients: calling Delete usecase")
	return nil
}
