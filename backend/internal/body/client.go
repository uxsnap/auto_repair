package body

import (
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/uxsnap/auto_repair/backend/internal/entity"
)

type CreateClientBody struct {
	Name         string
	EmployeeId   uuid.UUID
	Phone        string
	HasDocuments bool
	Passport     string
}

type ClientBodyParams struct {
	Name         string
	Phone        string
	Passport     string
	EmployeeName string
}

func (c *CreateClientBody) ToEntity() entity.Client {
	return entity.Client{
		Id: pgtype.UUID{
			Bytes:  uuid.New(),
			Status: pgtype.Present,
		},
		EmployeeId: pgtype.UUID{
			Bytes:  c.EmployeeId,
			Status: pgtype.Present,
		},
		Name:         c.Name,
		Phone:        c.Phone,
		HasDocuments: c.HasDocuments,
		Passport:     c.Passport,
	}
}
