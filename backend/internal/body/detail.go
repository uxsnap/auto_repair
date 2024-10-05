package body

import (
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/uxsnap/auto_repair/backend/internal/entity"
)

type CreateDetailBody struct {
	Name      string
	StorageId uuid.UUID
	Price     float64
}

func (c *CreateDetailBody) ToEntity() entity.Detail {
	return entity.Detail{
		Id: pgtype.UUID{
			Bytes:  uuid.New(),
			Status: pgtype.Present,
		},
		Name: c.Name,
		StorageId: pgtype.UUID{
			Bytes:  c.StorageId,
			Status: pgtype.Present,
		},
		Price: c.Price,
	}
}
