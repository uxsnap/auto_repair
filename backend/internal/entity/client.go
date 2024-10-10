package entity

import (
	"github.com/jackc/pgtype"
)

type Client struct {
	Id           pgtype.UUID `json:"id"`
	EmployeeId   pgtype.UUID `json:"employee_id"`
	Name         string      `json:"name"`
	Phone        string      `json:"phone"`
	HasDocuments bool        `json:"has_documents"`
	Passport     string      `json:"passport"`
	IsDeleted    bool        `json:"is_deleted"`
}
