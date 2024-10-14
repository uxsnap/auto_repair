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

type ActsRepository interface {
	GetAll(ctx context.Context, params body.ActBodyParams) ([]entity.ActWithData, error)
	Create(ctx context.Context, client entity.Act) (uuid.UUID, error)
	Delete(ctx context.Context, clientID string) (uuid.UUID, error)
}

type ActsService struct {
	repo ActsRepository
}

func NewActsService(repo ActsRepository) *ActsService {
	return &ActsService{
		repo: repo,
	}
}

func (cs *ActsService) GetAll(ctx context.Context, params body.ActBodyParams) ([]entity.ActWithData, error) {
	log.Println("Acts: calling GetAll usecase")

	return cs.repo.GetAll(ctx, params)
}

func (cs *ActsService) Create(ctx context.Context, clientData body.CreateActBody) (uuid.UUID, error) {
	log.Println("Acts: calling Create usecase")

	if !validators.IsValidGuid(clientData.ApplicationId) {
		return uuid.Nil, fmt.Errorf("ApplicationId должен быть UUID")
	}

	if !validators.IsValidGuid(clientData.ServiceId) {
		return uuid.Nil, fmt.Errorf("ServiceId должен быть UUID")
	}

	if !validators.IsValidLen(clientData.Name, 10) {
		return uuid.Nil, fmt.Errorf("длина имени должна быть больше 10 символов")
	}

	return cs.repo.Create(ctx, clientData.ToEntity())
}

func (cs *ActsService) Delete(ctx context.Context, clientID uuid.UUID) (uuid.UUID, error) {
	log.Println("Acts: calling Delete usecase")

	if clientID == uuid.Nil {
		return uuid.Nil, fmt.Errorf("id must be provided")
	}

	return cs.repo.Delete(ctx, clientID.String())
}
