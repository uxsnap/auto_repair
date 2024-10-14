package body

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/uxsnap/auto_repair/backend/internal/entity"
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
			Time:   time.Now(),
			Status: pgtype.Present,
		},
	}
}
