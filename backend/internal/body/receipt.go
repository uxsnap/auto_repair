package body

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/uxsnap/auto_repair/backend/internal/entity"
)

type CreateReceiptBody struct {
	Sum        int
	ContractId uuid.UUID
}

type ReceiptBodyParams struct {
	ContractName string
	MinSum       int
	MaxSum       int
	MinCreatedAt string
	MaxCreatedAt string
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
