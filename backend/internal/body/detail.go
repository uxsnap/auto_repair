package body

import (
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/uxsnap/auto_repair/backend/internal/entity"
)

type CreateDetailBody struct {
	Name  string
	Price int
	Type  string
}

type DetailBodyParams struct {
	Name     string
	MinPrice int
	MaxPrice int
	Type     string
}

func (c *CreateDetailBody) ToEntity() entity.Detail {
	return entity.Detail{
		Id: pgtype.UUID{
			Bytes:  uuid.New(),
			Status: pgtype.Present,
		},
		Name:  c.Name,
		Price: c.Price,
		Type:  c.Type,
	}
}
