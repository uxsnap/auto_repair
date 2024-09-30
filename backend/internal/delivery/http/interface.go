package handler

import (
	"context"

	"github.com/uxsnap/auto_repair/backend/internal/entity"
)

type ClientService interface {
	GetAllClients(ctx context.Context) ([]entity.Client, error)
}
