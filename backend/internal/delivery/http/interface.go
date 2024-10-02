package handler

import (
	"context"

	"github.com/uxsnap/auto_repair/backend/internal/entity"
)

type ClientsService interface {
	GetAll(ctx context.Context) ([]entity.Client, error)
	Create(ctx context.Context, clientData entity.CreateClientBody) error
}
