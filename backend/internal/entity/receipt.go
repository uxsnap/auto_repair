package entity

import (
	"github.com/jackc/pgtype"
)

type Receipt struct {
	Id         pgtype.UUID      `json:"id"`
	ContractId pgtype.UUID      `json:"contractId"`
	CreatedAt  pgtype.Timestamp `json:"created_at"`
	Sum        int              `json:"sum"`
	IsDeleted  bool             `json:"isDeleted"`
}
