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

	if !validators.IsValidLen(clientData.Name, 3) {
		return uuid.Nil, fmt.Errorf("длина имени должна быть больше 3 символов")
	}

	if !validators.IsValidSum(clientData.Sum) {
		return uuid.Nil, fmt.Errorf("сумма должна быть больше 0")
	}

	if !validators.IsValidGuid(clientData.StatusId) {
		return uuid.Nil, fmt.Errorf("неккоректный statusId")
	}

	return cs.repo.Create(ctx, clientData.ToEntity())
}

func (cs *ContractsService) Update(ctx context.Context, id uuid.UUID, clientData body.CreateContractBody) error {
	log.Println("Contracts: calling Update usecase")

	return cs.repo.Update(ctx, id, clientData.ToEntity())
}
