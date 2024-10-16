package body

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/uxsnap/auto_repair/backend/internal/entity"
	"github.com/uxsnap/auto_repair/backend/internal/utils"
)

type CreateActBody struct {
	Name          string
	ApplicationId uuid.UUID
	ServiceId     uuid.UUID
}

type ActBodyParams struct {
	Name            string
	ApplicationName string
	ServiceName     string
	MinCreatedAt    string
	MaxCreatedAt    string
}

func (c *CreateActBody) ToEntity() entity.Act {
	return entity.Act{
		Id: pgtype.UUID{
			Bytes:  uuid.New(),
			Status: pgtype.Present,
		},
		Name: c.Name,
		ApplicationId: pgtype.UUID{
			Bytes:  c.ApplicationId,
			Status: pgtype.Present,
		},
		ServiceId: pgtype.UUID{
			Bytes:  c.ServiceId,
			Status: pgtype.Present,
		},
		CreatedAt: pgtype.Timestamp{
			Time:   utils.TruncTime(time.Now()),
			Status: pgtype.Present,
		},
	}
}
