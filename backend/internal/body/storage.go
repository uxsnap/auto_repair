package body

import (
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/uxsnap/auto_repair/backend/internal/entity"
)

type CreateStorageBody struct {
	EmployeeId  uuid.UUID
	DetailId    uuid.UUID
	StorageNum  string
	DetailCount int
}

type StorageBodyParams struct {
	StorageNum   string
	DetailName   string
	EmployeeName string
}

func (c *CreateStorageBody) ToEntity() entity.Storage {
	return entity.Storage{
		Id: pgtype.UUID{
			Bytes:  uuid.New(),
			Status: pgtype.Present,
		},
		EmployeeId: pgtype.UUID{
			Bytes:  c.EmployeeId,
			Status: pgtype.Present,
		},
		StorageNum: c.StorageNum,
		DetailId: pgtype.UUID{
			Bytes:  c.DetailId,
			Status: pgtype.Present,
		},
		DetailCount: c.DetailCount,
	}
}
