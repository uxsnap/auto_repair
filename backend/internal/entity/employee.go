package entity

import "github.com/jackc/pgtype"

type Employee struct {
	Id          pgtype.UUID `json:"id"`
	Name        string      `json:"name"`
	Position    string      `json:"position"`
	EmployeeNum string      `json:"employeeNum"`
}
