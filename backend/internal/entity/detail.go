package entity

import (
	"github.com/jackc/pgtype"
)

type Detail struct {
	Id    pgtype.UUID `json:"id"`
	Name  string      `json:"name"`
	Price int         `json:"price"`
	Type  string      `json:"type"`
}
