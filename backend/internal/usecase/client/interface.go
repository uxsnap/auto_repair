package usecase

import "github.com/uxsnap/auto_repair/backend/internal/entity"

type ClientRepository interface {
	GetAllClients() ([]entity.Client, error)
}
