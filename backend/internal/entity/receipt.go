package entity

import (
	"github.com/jackc/pgtype"
)

type Receipt struct {
	Id         pgtype.UUID      `json:"id"`
	ContractId pgtype.UUID      `json:"contractId"`
	CreatedAt  pgtype.Timestamp `json:"createdAt"`
	Sum        int              `json:"sum"`
	IsDeleted  bool             `json:"isDeleted"`
}

type ReceiptWithData struct {
	Id        pgtype.UUID      `json:"id"`
	Contract  IdWithName       `json:"contract"`
	CreatedAt pgtype.Timestamp `json:"createdAt"`
	Sum       int              `json:"sum"`
}
