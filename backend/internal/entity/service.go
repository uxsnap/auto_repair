package entity

import (
	"github.com/jackc/pgtype"
)

type Service struct {
	Id   pgtype.UUID `json:"id"`
	Name string      `json:"name"`
}
