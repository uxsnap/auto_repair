package body

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/uxsnap/auto_repair/backend/internal/entity"
)

type CreateContractBody struct {
	Name      string
	Sum       int
	CreatedAt time.Time
	SignedAt  time.Time
	StatusId  uuid.UUID
}

func (c *CreateContractBody) ToEntity() entity.Contract {
	return entity.Contract{
		Id: pgtype.UUID{
			Bytes:  uuid.New(),
			Status: pgtype.Present,
		},
		Sum: c.Sum,
		CreatedAt: pgtype.Timestamp{
			Time:   c.CreatedAt,
			Status: pgtype.Present,
		},
		SignedAt: pgtype.Timestamp{
			Time:   c.CreatedAt,
			Status: pgtype.Present,
		},
		StatusId: pgtype.UUID{
			Bytes:  c.StatusId,
			Status: pgtype.Present,
		},
	}
}
