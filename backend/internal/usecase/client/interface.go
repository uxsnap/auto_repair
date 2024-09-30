package useCaseClients

import "github.com/uxsnap/auto_repair/backend/internal/entity"

type ClientsRepository interface {
	GetAllClients() ([]entity.Client, error)
}
