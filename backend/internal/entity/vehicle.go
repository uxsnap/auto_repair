package entity

import "github.com/jackc/pgtype"

type Vehicle struct {
	Id            pgtype.UUID `json:"id"`
	ClientId      pgtype.UUID `json:"clientId"`
	VehicleNumber string      `json:"vehicleNumber"`
	Brand         string      `json:"brand"`
	Model         string      `json:"model"`
	IsDeleted     bool        `json:"isDeleted"`
}

type VehicleWithData struct {
	Id            pgtype.UUID `json:"id"`
	Client        IdWithName  `json:"client"`
	VehicleNumber string      `json:"vehicleNumber"`
	Brand         string      `json:"brand"`
	Model         string      `json:"model"`
	IsDeleted     bool        `json:"isDeleted"`
}
