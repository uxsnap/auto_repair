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
	GetAll(ctx context.Context, params body.VehicleBodyParams) ([]entity.VehicleWithData, error)
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

func (cs *VehiclesService) GetAll(ctx context.Context, params body.VehicleBodyParams) ([]entity.VehicleWithData, error) {
	log.Println("Vehicles: calling GetAll usecase")

	return cs.repo.GetAll(ctx, params)
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

	if !validators.IsValidLenEq(clientData.VehicleNumber, 8) {
		return uuid.Nil, fmt.Errorf("длина номера должна быть равна 8 символам")
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
