package entity

import "github.com/jackc/pgtype"

type Vehicle struct {
	Id            pgtype.UUID `json:"id"`
	ClientId      pgtype.UUID `json:"clientId"`
	VehicleNumber string      `json:"vehicleNumber"`
	Brand         string      `json:"brand"`
	Model         string      `json:"Model"`
	IsDeleted     bool        `json:"isDeleted"`
}
