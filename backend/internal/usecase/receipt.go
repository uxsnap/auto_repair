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

type ReceiptsRepository interface {
	GetAll(ctx context.Context, params body.ReceiptBodyParams) ([]entity.ReceiptWithData, error)
	Create(ctx context.Context, Receipt entity.Receipt) (uuid.UUID, error)
}

type ReceiptsService struct {
	repo ReceiptsRepository
}

func NewReceiptsService(repo ReceiptsRepository) *ReceiptsService {
	return &ReceiptsService{
		repo: repo,
	}
}

func (cs *ReceiptsService) GetAll(ctx context.Context, params body.ReceiptBodyParams) ([]entity.ReceiptWithData, error) {
	log.Println("Receipts: calling GetAll usecase")

	return cs.repo.GetAll(ctx, params)
}

func (cs *ReceiptsService) Create(ctx context.Context, clientData body.CreateReceiptBody) (uuid.UUID, error) {
	log.Println("Receipts: calling Create usecase")

	if !validators.IsValidSum(clientData.Sum) {
		return uuid.Nil, fmt.Errorf("сумма должна быть больше 0")
	}

	if !validators.IsValidGuid(clientData.ContractId) {
		return uuid.Nil, fmt.Errorf("неккоректный ContractId")
	}

	return cs.repo.Create(ctx, clientData.ToEntity())
}
