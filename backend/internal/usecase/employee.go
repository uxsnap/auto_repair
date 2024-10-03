package usecase

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/uxsnap/auto_repair/backend/internal/body"
	"github.com/uxsnap/auto_repair/backend/internal/entity"
)

type EmployeesRepository interface {
	GetAll(ctx context.Context) ([]entity.Employee, error)
	Create(ctx context.Context, Employee entity.Employee) (uuid.UUID, error)
	Update(ctx context.Context, id uuid.UUID, Employee entity.Employee) error
}

type EmployeesService struct {
	repo EmployeesRepository
}

func NewEmployeesService(repo EmployeesRepository) *EmployeesService {
	return &EmployeesService{
		repo: repo,
	}
}

func (cs *EmployeesService) GetAll(ctx context.Context) ([]entity.Employee, error) {
	log.Println("Employees: calling GetAll usecase")

	return cs.repo.GetAll(ctx)
}

func (cs *EmployeesService) Create(ctx context.Context, clientData body.CreateEmployeeBody) (uuid.UUID, error) {
	log.Println("Employees: calling Create usecase")

	if len(clientData.Name) < 3 {
		return uuid.Nil, fmt.Errorf("длина имени должна быть больше 3 символов")
	}

	if len(clientData.EmployeeNum) < 5 {
		return uuid.Nil, fmt.Errorf("неверный формат номера сотрудника")
	}

	if len(clientData.Position) < 5 {
		return uuid.Nil, fmt.Errorf("неверный формат должности сотрудника")
	}

	return cs.repo.Create(ctx, clientData.ToEntity())
}

func (cs *EmployeesService) Update(ctx context.Context, id uuid.UUID, clientData body.CreateEmployeeBody) error {
	log.Println("Employees: calling Update usecase")

	if len(clientData.Name) < 3 {
		return fmt.Errorf("длина имени должна быть больше 3 символов")
	}

	if len(clientData.EmployeeNum) < 5 {
		return fmt.Errorf("неверный формат номера сотрудника")
	}

	if len(clientData.Position) < 5 {
		return fmt.Errorf("неверный формат должности сотрудника")
	}

	return cs.repo.Update(ctx, id, clientData.ToEntity())
}
