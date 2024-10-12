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
	GetAll(ctx context.Context, params body.StorageBodyParams) ([]entity.StorageWithData, error)
	Create(ctx context.Context, client entity.Storage) (uuid.UUID, error)
	Delete(ctx context.Context, storageID string) (uuid.UUID, error)
	Update(ctx context.Context, id uuid.UUID, clientData body.CreateStorageBody) error
}

type StoragesService struct {
	repo StoragesRepository
}

func NewStoragesService(repo StoragesRepository) *StoragesService {
	return &StoragesService{
		repo: repo,
	}
}

func (cs *StoragesService) GetAll(ctx context.Context, params body.StorageBodyParams) ([]entity.StorageWithData, error) {
	log.Println("Storages: calling GetAll usecase")

	return cs.repo.GetAll(ctx, params)
}

func (cs *StoragesService) Create(ctx context.Context, clientData body.CreateStorageBody) (uuid.UUID, error) {
	log.Println("Storages: calling Create usecase")

	if !validators.IsValidGuid(clientData.EmployeeId) {
		return uuid.Nil, fmt.Errorf("ApplicationId должен быть UUID")
	}

	if !validators.IsValidLen(clientData.StorageNum, 3) {
		return uuid.Nil, fmt.Errorf("длина номера должна быть больше 3 символов")
	}

	if !validators.IsValidSum(clientData.DetailCount) {
		return uuid.Nil, fmt.Errorf("кол-во деталей не должно быть равно нулю")
	}

	return cs.repo.Create(ctx, clientData.ToEntity())
}

func (cs *StoragesService) Update(ctx context.Context, id uuid.UUID, clientData body.CreateStorageBody) error {
	log.Println("Storages: calling Update usecase")

	if !validators.IsValidGuid(clientData.EmployeeId) {
		return fmt.Errorf("ApplicationId должен быть UUID")
	}

	if !validators.IsValidLen(clientData.StorageNum, 3) {
		return fmt.Errorf("длина номера должна быть больше 3 символов")
	}

	if !validators.IsValidSum(clientData.DetailCount) {
		return fmt.Errorf("кол-во деталей не должно быть равно нулю")
	}

	return cs.repo.Update(ctx, id, clientData)
}

func (cs *StoragesService) Delete(ctx context.Context, detailID uuid.UUID) (uuid.UUID, error) {
	log.Println("Storage: calling Delete usecase")

	if detailID == uuid.Nil {
		return uuid.Nil, fmt.Errorf("id must be provided")
	}

	return cs.repo.Delete(ctx, detailID.String())
}
