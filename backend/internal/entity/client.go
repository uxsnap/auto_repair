package entity

import (
	"github.com/gofrs/uuid"
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
