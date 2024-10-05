package entity

import (
	"github.com/jackc/pgtype"
)

type Act struct {
	Id            pgtype.UUID      `json:"id"`
	Name          string           `json:"name"`
	CreatedAt     pgtype.Timestamp `json:"created_at"`
	ApplicationId pgtype.UUID      `json:"application_id"`
	ServiceId     pgtype.UUID      `json:"service_id"`
	IsDeleted     bool
}
