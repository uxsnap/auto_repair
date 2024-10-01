package handler

import (
	"context"

	"github.com/uxsnap/auto_repair/backend/internal/entity"
)

type ClientsService interface {
	GetAllClients(ctx context.Context) ([]entity.Client, error)
}
