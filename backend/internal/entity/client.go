package entity

import (
	"github.com/jackc/pgtype"
)

type Client struct {
	Id           pgtype.UUID `json:"id"`
	EmployeeId   pgtype.UUID `json:"employeeId"`
	Name         string      `json:"name"`
	Phone        string      `json:"phone"`
	HasDocuments bool        `json:"hasDocuments"`
	Passport     string      `json:"passport"`
	IsDeleted    bool        `json:"isDeleted"`
}

func (c Client) Table() string {
	return "clients"
}

func (c Client) Columns() string {
	return "name"
	// "id, employee_id, name, phone, passport, has_documents, is_deleted"
}
