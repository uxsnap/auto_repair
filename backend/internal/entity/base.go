package entity

import "github.com/jackc/pgtype"

type IdWithName struct {
	Id   pgtype.UUID `json:"id"`
	Name string      `json:"name"`
}
