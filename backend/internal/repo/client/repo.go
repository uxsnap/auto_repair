package repoClients

import (
	"github.com/uxsnap/auto_repair/backend/internal/db"
	"github.com/uxsnap/auto_repair/backend/internal/entity"
	"github.com/uxsnap/auto_repair/backend/internal/repo"
)

type ClientsRepository struct {
	*repo.BasePgRepository
}

func NewClientsRepo(client *db.Client) *ClientsRepository {
	return &ClientsRepository{
		repo.NewBaseRepo(client),
	}
}

func (cr *ClientsRepository) GetAllClients() ([]entity.Client, error) {
	var clients []entity.Client

	clients = append(clients, entity.Client{
		Name: "test",
	})

	return clients, nil
}
