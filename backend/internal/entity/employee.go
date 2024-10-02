package entity

import "github.com/jackc/pgtype"

type Employee struct {
	Id           pgtype.UUID `json:"id"`
	Name         string      `json:"name"`
	Position     string      `json:"position"`
	HasDocuments bool        `json:"hasDocuments"`
	EmployeeNum  string      `json:"employeeNum"`
	IsDeleted    bool        `json:"isDeleted"`
}
