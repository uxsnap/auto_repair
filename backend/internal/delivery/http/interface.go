package handler

import (
	"github.com/uxsnap/auto_repair/backend/internal/entity"
)

type ClientsService interface {
	GetAllClients() ([]entity.Client, error)
}
