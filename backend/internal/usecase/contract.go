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
	GetAll(ctx context.Context, params body.ContractBodyParams) ([]entity.Contract, error)
	Create(ctx context.Context, client entity.Contract) (uuid.UUID, error)
	Update(ctx context.Context, id uuid.UUID, client entity.Contract) error
	Delete(ctx context.Context, ID string) (uuid.UUID, error)
}

type ContractsService struct {
	repo ContractsRepository
}

func NewContractsService(repo ContractsRepository) *ContractsService {
	return &ContractsService{
		repo: repo,
	}
}

func (cs *ContractsService) GetAll(ctx context.Context, params body.ContractBodyParams) ([]entity.Contract, error) {
	log.Println("Contracts: calling GetAll usecase")

	return cs.repo.GetAll(ctx, params)
}

func (cs *ContractsService) Create(ctx context.Context, clientData body.CreateContractBody) (uuid.UUID, error) {
	log.Println("Contracts: calling Create usecase")

	if !validators.IsValidLen(clientData.Name, 3) {
		return uuid.Nil, fmt.Errorf("длина имени должна быть больше 3 символов")
	}

	if !validators.IsValidSum(clientData.Sum) {
		return uuid.Nil, fmt.Errorf("сумма должна быть больше 0")
	}

	if !validators.IsValidLen(clientData.Status, 0) {
		return uuid.Nil, fmt.Errorf("неккоректный statusId")
	}

	return cs.repo.Create(ctx, clientData.ToEntity())
}

func (cs *ContractsService) Update(ctx context.Context, id uuid.UUID, clientData body.CreateContractBody) error {
	log.Println("Contracts: calling Update usecase")

	return cs.repo.Update(ctx, id, clientData.ToEntity())
}

func (cs *ContractsService) Delete(ctx context.Context, clientID uuid.UUID) (uuid.UUID, error) {
	log.Println("contracts: calling Delete usecase")

	if clientID == uuid.Nil {
		return uuid.Nil, fmt.Errorf("id must be provided")
	}

	return cs.repo.Delete(ctx, clientID.String())
}
