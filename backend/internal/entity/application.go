package entity

import (
	"github.com/jackc/pgtype"
)

type Application struct {
	Id         pgtype.UUID      `json:"id"`
	EmployeeId pgtype.UUID      `json:"employeeId"`
	ClientId   pgtype.UUID      `json:"clientId"`
	Name       string           `json:"name"`
	CreatedAt  pgtype.Timestamp `json:"createdAt"`
	Status     string           `json:"status"`
	ContractId pgtype.UUID      `json:"contractId"`
	IsDeleted  bool             `json:"isDeleted"`
}
