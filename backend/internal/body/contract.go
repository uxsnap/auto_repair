package body

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/uxsnap/auto_repair/backend/internal/entity"
	"github.com/uxsnap/auto_repair/backend/internal/utils"
)

type CreateContractBody struct {
	Name     string
	Sum      int
	SignedAt time.Time
	Status   string
}

type ContractBodyParams struct {
	Name         string
	MinSum       int
	MaxSum       int
	Status       string
	MinCreatedAt string
	MaxCreatedAt string
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
			Time:   utils.TruncTime(time.Now()),
			Status: pgtype.Present,
		},
		SignedAt: pgtype.Timestamp{
			Time:   utils.TruncTime(c.SignedAt),
			Status: pgtype.Present,
		},
		Status: c.Status,
	}
}
