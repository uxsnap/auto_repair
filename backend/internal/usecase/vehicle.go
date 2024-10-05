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

type VehiclesRepository interface {
	GetAll(ctx context.Context) ([]entity.Vehicle, error)
	Create(ctx context.Context, client entity.Vehicle) (uuid.UUID, error)
	Delete(ctx context.Context, clientID string) (uuid.UUID, error)
}

type VehiclesService struct {
	repo VehiclesRepository
}

func NewVehiclesService(repo VehiclesRepository) *VehiclesService {
	return &VehiclesService{
		repo: repo,
	}
}

func (cs *VehiclesService) GetAll(ctx context.Context) ([]entity.Vehicle, error) {
	log.Println("Vehicles: calling GetAll usecase")

	return cs.repo.GetAll(ctx)
}

func (cs *VehiclesService) Create(ctx context.Context, clientData body.CreateVehicleBody) (uuid.UUID, error) {
	log.Println("Vehicles: calling Create usecase")

	if !validators.IsValidGuid(clientData.ClientId) {
		return uuid.Nil, fmt.Errorf("clientId должен быть UUID")
	}

	if !validators.IsValidLen(clientData.Brand, 3) {
		return uuid.Nil, fmt.Errorf("длина бренда должна быть больше 3 символов")
	}

	if !validators.IsValidLen(clientData.Model, 3) {
		return uuid.Nil, fmt.Errorf("длина модели должна быть больше 3 символов")
	}

	if !validators.IsValidLenEq(clientData.VehicleNumber, 10) {
		return uuid.Nil, fmt.Errorf("длина номера должна быть равна 10 символов")
	}

	if !validators.IsValidPass(clientData.VehicleNumber) {
		return uuid.Nil, fmt.Errorf("неверный формат паспорта")
	}

	return cs.repo.Create(ctx, clientData.ToEntity())
}

func (cs *VehiclesService) Delete(ctx context.Context, clientID uuid.UUID) (uuid.UUID, error) {
	log.Println("Vehicles: calling Delete usecase")

	if clientID == uuid.Nil {
		return uuid.Nil, fmt.Errorf("id не найден")
	}

	return cs.repo.Delete(ctx, clientID.String())
}
