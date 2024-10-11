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

type DetailsRepository interface {
	GetAll(ctx context.Context, params body.DetailBodyParams) ([]entity.Detail, error)
	Create(ctx context.Context, detail entity.Detail) (uuid.UUID, error)
	Delete(ctx context.Context, detailID string) (uuid.UUID, error)
}

type DetailsService struct {
	repo DetailsRepository
}

func NewDetailsService(repo DetailsRepository) *DetailsService {
	return &DetailsService{
		repo: repo,
	}
}

func (cs *DetailsService) GetAll(ctx context.Context, params body.DetailBodyParams) ([]entity.Detail, error) {
	log.Println("Details: calling GetAll usecase")

	return cs.repo.GetAll(ctx, params)
}

func (cs *DetailsService) Create(ctx context.Context, detailData body.CreateDetailBody) (uuid.UUID, error) {
	log.Println("Details: calling Create usecase")

	if !validators.IsValidLen(detailData.Name, 3) {
		return uuid.Nil, fmt.Errorf("длина имени должна быть больше 3 символов")
	}

	if !validators.IsValidLen(detailData.Type, 3) {
		return uuid.Nil, fmt.Errorf("длина типа детали должна быть больше 3 символов")
	}

	if !validators.IsValidSum(detailData.Price) {
		return uuid.Nil, fmt.Errorf("цена должна быть больше нуля")
	}

	return cs.repo.Create(ctx, detailData.ToEntity())
}

func (cs *DetailsService) Delete(ctx context.Context, detailID uuid.UUID) (uuid.UUID, error) {
	log.Println("Details: calling Delete usecase")

	if detailID == uuid.Nil {
		return uuid.Nil, fmt.Errorf("id must be provided")
	}

	return cs.repo.Delete(ctx, detailID.String())
}
