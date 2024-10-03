package body

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/uxsnap/auto_repair/backend/internal/entity"
)

type CreateContractBody struct {
	Name     string
	Sum      int
	SignedAt time.Time
	StatusId uuid.UUID
}

func (c *CreateContractBody) ToEntity() entity.Contract {
	return entity.Contract{
		Id: pgtype.UUID{
			Bytes:  uuid.New(),
			Status: pgtype.Present,
		},
		Name: c.Name,
		Sum:  c.Sum,
		CreatedAt: pgtype.Timestamp{
			Time:   time.Now(),
			Status: pgtype.Present,
		},
		StatusId: pgtype.UUID{
			Bytes:  c.StatusId,
			Status: pgtype.Present,
		},
	}
}
