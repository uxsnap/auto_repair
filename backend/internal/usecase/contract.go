package usecase

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/uxsnap/auto_repair/backend/internal/body"
	"github.com/uxsnap/auto_repair/backend/internal/entity"
)

type ContractsRepository interface {
	GetAll(ctx context.Context) ([]entity.Contract, error)
	Create(ctx context.Context, client entity.Contract) (uuid.UUID, error)
	Update(ctx context.Context, id uuid.UUID, client entity.Contract) error
}

type ContractsService struct {
	repo ContractsRepository
}

func NewContractsService(repo ContractsRepository) *ContractsService {
	return &ContractsService{
		repo: repo,
	}
}

func (cs *ContractsService) GetAll(ctx context.Context) ([]entity.Contract, error) {
	log.Println("Contracts: calling GetAll usecase")

	return cs.repo.GetAll(ctx)
}

func (cs *ContractsService) Create(ctx context.Context, clientData body.CreateContractBody) (uuid.UUID, error) {
	log.Println("Contracts: calling Create usecase")

	// if err := uuid.Validate(clientData.EmployeeId.String()); err != nil || clientData.EmployeeId == uuid.Nil {
	// 	return uuid.Nil, fmt.Errorf("employeeId должен быть UUID")
	// }

	// if len(clientData.Name) < 3 {
	// 	return uuid.Nil, fmt.Errorf("длина имени должна быть больше 3 символов")
	// }

	// if !IsValidPhoneNumber(clientData.Phone) {
	// 	return uuid.Nil, fmt.Errorf("неверный формат номера")
	// }

	// if !IsValidPass(clientData.Passport) {
	// 	return uuid.Nil, fmt.Errorf("неверный формат паспорта")
	// }

	return cs.repo.Create(ctx, clientData.ToEntity())
}

func (cs *ContractsService) Update(ctx context.Context, id uuid.UUID, clientData body.CreateContractBody) error {
	log.Println("Contracts: calling Update usecase")

	return cs.repo.Update(ctx, id, clientData.ToEntity())
}
