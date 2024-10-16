package entity

import (
	"github.com/jackc/pgtype"
)

type Act struct {
	Id            pgtype.UUID      `json:"id"`
	Name          string           `json:"name"`
	CreatedAt     pgtype.Timestamp `json:"createdAt"`
	ApplicationId pgtype.UUID      `json:"applicationId"`
	ServiceId     pgtype.UUID      `json:"serviceId"`
	IsDeleted     bool
}

type ActWithData struct {
	Id          pgtype.UUID      `json:"id"`
	Name        string           `json:"name"`
	CreatedAt   pgtype.Timestamp `json:"createdAt"`
	Application IdWithName       `json:"application"`
	Service     IdWithName       `json:"service"`
}
