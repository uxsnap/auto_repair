package useCaseClients

import "github.com/uxsnap/auto_repair/backend/internal/entity"

type ClientsService struct {
	repo ClientsRepository
}

func NewClientsService(repo ClientsRepository) *ClientsService {
	return &ClientsService{
		repo: repo,
	}
}

func (cs *ClientsService) GetAllClients() ([]entity.Client, error) {
	return cs.repo.GetAllClients()
}
