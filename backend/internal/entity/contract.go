package entity

import (
	"github.com/jackc/pgtype"
)

type Contract struct {
	Id        pgtype.UUID      `json:"id"`
	Name      string           `json:"name"`
	Sum       int              `json:"sum"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	SignedAt  pgtype.Timestamp `json:"signed_at"`
	StatusId  pgtype.UUID      `json:"status_id"`
}
