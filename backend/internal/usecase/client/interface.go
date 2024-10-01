package useCaseClients

import (
	"context"

	"github.com/uxsnap/auto_repair/backend/internal/entity"
)

type ClientsRepository interface {
	GetAll(ctx context.Context) ([]entity.Client, error)
	Create(ctx context.Context, client entity.Client) error
	Delete(ctx context.Context, clientID string) error
	Update(ctx context.Context, clientData entity.Client) error
}
