package useCaseClients

import (
	"context"

	"github.com/uxsnap/auto_repair/backend/internal/entity"
)

type ClientsRepository interface {
	GetAllClients(ctx context.Context) ([]entity.Client, error)
}
