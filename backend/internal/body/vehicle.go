package body

import (
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/uxsnap/auto_repair/backend/internal/entity"
)

type CreateVehicleBody struct {
	ClientId      uuid.UUID
	VehicleNumber string
	Brand         string
	Model         string
}

type VehicleBodyParams struct {
	VehicleNumber string
	Brand         string
	Model         string
}

func (c *CreateVehicleBody) ToEntity() entity.Vehicle {
	return entity.Vehicle{
		Id: pgtype.UUID{
			Bytes:  uuid.New(),
			Status: pgtype.Present,
		},
		ClientId: pgtype.UUID{
			Bytes:  c.ClientId,
			Status: pgtype.Present,
		},
		Brand:         c.Brand,
		Model:         c.Model,
		VehicleNumber: c.VehicleNumber,
	}
}
