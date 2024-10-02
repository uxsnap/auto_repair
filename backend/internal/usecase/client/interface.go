package useCaseClients

import (
	"context"

	"github.com/google/uuid"
	"github.com/uxsnap/auto_repair/backend/internal/entity"
	"github.com/uxsnap/auto_repair/backend/internal/repo"
)

type ClientsRepository interface {
	GetAll(ctx context.Context, baseEntity repo.BaseEntity, dst interface{}) error
	Create(ctx context.Context, client entity.Client) (uuid.UUID, error)
	Delete(ctx context.Context, clientID string) (uuid.UUID, error)
	Update(ctx context.Context, id uuid.UUID, client entity.Client) error
}
