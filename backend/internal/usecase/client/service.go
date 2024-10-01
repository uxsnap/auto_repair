package useCaseClients

import (
	"context"
	"log"

	"github.com/jackc/pgtype"
	"github.com/uxsnap/auto_repair/backend/internal/entity"
)

type ClientsService struct {
	repo ClientsRepository
}

func NewClientsService(repo ClientsRepository) *ClientsService {
	return &ClientsService{
		repo: repo,
	}
}

func (cs *ClientsService) GetAll(ctx context.Context) ([]entity.Client, error) {
	log.Println("clients: calling GetAll usecase")
	return cs.repo.GetAll(ctx)
}

func (cs *ClientsService) Create(ctx context.Context, clientData entity.Client) error {
	log.Println("clients: calling Create usecase")
	return cs.repo.Create(ctx, clientData)
}

func (cs *ClientsService) Update(ctx context.Context, clientData entity.Client) error {
	log.Println("clients: calling Update usecase")
	return nil
}

func (cs *ClientsService) Delete(ctx context.Context, clientID pgtype.UUID) error {
	log.Println("clients: calling Delete usecase")
	return nil
}
