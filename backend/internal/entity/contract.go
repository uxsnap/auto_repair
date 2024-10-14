package entity

import (
	"github.com/jackc/pgtype"
)

type Contract struct {
	Id        pgtype.UUID      `json:"id"`
	Name      string           `json:"name"`
	Sum       int              `json:"sum"`
	CreatedAt pgtype.Timestamp `json:"createdAt"`
	SignedAt  pgtype.Timestamp `json:"signedAt"`
	Status    string           `json:"status"`
}
