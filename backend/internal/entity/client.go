package entity

import (
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

type CreateClientBody struct {
	Name         string
	EmployeeId   uuid.UUID
	Phone        string
	HasDocuments bool
	Passport     string
}

type Client struct {
	Id           pgtype.UUID `json:"id"`
	EmployeeId   pgtype.UUID `json:"employeeId"`
	Name         string      `json:"name"`
	Phone        string      `json:"phone"`
	HasDocuments bool        `json:"hasDocuments"`
	Passport     string      `json:"passport"`
}

func (c *CreateClientBody) ToEntity() Client {
	return Client{
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
