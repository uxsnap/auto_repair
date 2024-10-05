package entity

import (
	"github.com/jackc/pgtype"
)

type Application struct {
	Id         pgtype.UUID      `json:"id"`
	EmployeeId pgtype.UUID      `json:"employee_id"`
	ClientId   pgtype.UUID      `json:"client_id"`
	Name       string           `json:"name"`
	CreatedAt  pgtype.Timestamp `json:"created_at"`
	Status     string           `json:"status"`
	ContractId pgtype.UUID      `json:"contract_id"`
	IsDeleted  bool             `json:"is_deleted"`
}
