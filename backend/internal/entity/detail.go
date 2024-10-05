package entity

import (
	"github.com/jackc/pgtype"
)

type Detail struct {
	Id        pgtype.UUID `json:"id"`
	StorageId pgtype.UUID `json:"storage_id"`
	Name      string      `json:"name"`
	Price     float64     `json:"price"`
}
