package body

import (
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/uxsnap/auto_repair/backend/internal/entity"
)

type CreateEmployeeBody struct {
	Name        string
	Position    string
	EmployeeNum string
}

func (c *CreateEmployeeBody) ToEntity() entity.Employee {
	return entity.Employee{
		Id: pgtype.UUID{
			Bytes:  uuid.New(),
			Status: pgtype.Present,
		},
		Name:        c.Name,
		Position:    c.Position,
		EmployeeNum: c.EmployeeNum,
	}
}
