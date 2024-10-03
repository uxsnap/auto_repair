package body

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/uxsnap/auto_repair/backend/internal/entity"
)

type CreateReceiptBody struct {
	Sum        int
	CreatedAt  time.Time
	ContractId uuid.UUID
}

func (c *CreateReceiptBody) ToEntity() entity.Receipt {
	return entity.Receipt{
		Id: pgtype.UUID{
			Bytes:  uuid.New(),
			Status: pgtype.Present,
		},
		Sum: c.Sum,
		CreatedAt: pgtype.Timestamp{
			Time:   time.Now(),
			Status: pgtype.Present,
		},
		ContractId: pgtype.UUID{
			Bytes:  c.ContractId,
			Status: pgtype.Present,
		},
	}
}
