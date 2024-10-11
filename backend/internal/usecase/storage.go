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

type StoragesRepository interface {
	GetAll(ctx context.Context) ([]entity.Storage, error)
	Create(ctx context.Context, client entity.Storage) (uuid.UUID, error)
	Delete(ctx context.Context, storageID string) (uuid.UUID, error)
}

type StoragesService struct {
	repo StoragesRepository
}

func NewStoragesService(repo StoragesRepository) *StoragesService {
	return &StoragesService{
		repo: repo,
	}
}

func (cs *StoragesService) GetAll(ctx context.Context) ([]entity.Storage, error) {
	log.Println("Storages: calling GetAll usecase")

	return cs.repo.GetAll(ctx)
}

func (cs *StoragesService) Create(ctx context.Context, clientData body.CreateStorageBody) (uuid.UUID, error) {
	log.Println("Storages: calling Create usecase")

	if !validators.IsValidGuid(clientData.EmployeeId) {
		return uuid.Nil, fmt.Errorf("ApplicationId должен быть UUID")
	}

	if !validators.IsValidLen(clientData.StorageNum, 10) {
		return uuid.Nil, fmt.Errorf("длина имени должна быть больше 10 символов")
	}

	return cs.repo.Create(ctx, clientData.ToEntity())
}

func (cs *StoragesService) Delete(ctx context.Context, detailID uuid.UUID) (uuid.UUID, error) {
	log.Println("Storage: calling Delete usecase")

	if detailID == uuid.Nil {
		return uuid.Nil, fmt.Errorf("id must be provided")
	}

	return cs.repo.Delete(ctx, detailID.String())
}
