package handler

import (
	"context"

	"github.com/google/uuid"
	"github.com/uxsnap/auto_repair/backend/internal/body"
	"github.com/uxsnap/auto_repair/backend/internal/entity"
)

type ClientsService interface {
	GetAll(ctx context.Context) ([]entity.Client, error)
	Create(ctx context.Context, clientData body.CreateClientBody) (uuid.UUID, error)
	Update(ctx context.Context, cliendID uuid.UUID, clientData body.CreateClientBody) error
	Delete(ctx context.Context, cliendID uuid.UUID) (uuid.UUID, error)
}
