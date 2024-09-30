package repo

import (
	"github.com/uxsnap/auto_repair/backend/internal/db"
	"github.com/uxsnap/auto_repair/backend/internal/repo"
)

type ClientRepository struct {
	*repo.BasePgRepository
}

func NewClientRepo(client *db.Client) *ClientRepository {
	return &ClientRepository{
		repo.NewBaseRepo(client),
	}
}
