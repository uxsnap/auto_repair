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
	SignedAt string
	Status   string
}

type ContractBodyParams struct {
	Name   string
	MinSum int
	MaxSum int
	Status string
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
		Status: c.Status,
	}
}
