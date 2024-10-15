package body

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/uxsnap/auto_repair/backend/internal/entity"
)

type CreateApplicationBody struct {
	EmployeeId uuid.UUID
	ClientId   uuid.UUID
	Name       string
	Status     string
	ContractId uuid.UUID
}

type ApplicationBodyParams struct {
	Name         string
	ClientName   string
	EmployeeName string
	ContractName string
	Status       string
	// CreatedAt  time.Time
}

func (c *CreateApplicationBody) ToEntity() entity.Application {
	return entity.Application{
		Id: pgtype.UUID{
			Bytes:  uuid.New(),
			Status: pgtype.Present,
		},
		EmployeeId: pgtype.UUID{
			Bytes:  c.EmployeeId,
			Status: pgtype.Present,
		},
		ClientId: pgtype.UUID{
			Bytes:  c.ClientId,
			Status: pgtype.Present,
		},
		Name:   c.Name,
		Status: c.Status,
		ContractId: pgtype.UUID{
			Bytes:  c.ContractId,
			Status: pgtype.Present,
		},
		CreatedAt: pgtype.Timestamp{
			Time:   time.Now(),
			Status: pgtype.Present,
		},
	}
}
