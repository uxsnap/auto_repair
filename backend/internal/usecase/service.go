package usecase

import (
	"context"
	"log"

	"github.com/uxsnap/auto_repair/backend/internal/entity"
)

type ServicesRepository interface {
	GetAll(ctx context.Context) ([]entity.Service, error)
}

type ServicesService struct {
	repo ServicesRepository
}

func NewServicesService(repo ServicesRepository) *ServicesService {
	return &ServicesService{
		repo: repo,
	}
}

func (cs *ServicesService) GetAll(ctx context.Context) ([]entity.Service, error) {
	log.Println("Services: calling GetAll usecase")

	return cs.repo.GetAll(ctx)
}
