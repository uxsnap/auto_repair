package repoClients

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
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

func (cr *ClientsRepository) GetAllClients(ctx context.Context) ([]entity.Client, error) {
	log.Println("calling GetAllClients from repo")

	clientsSql := sq.Select("*").From("clients")

	sql, _, err := clientsSql.ToSql()

	if err != nil {
		log.Println("GetAllClients has errored")
		return nil, err
	}

	clients := []entity.Client{}

	pgxscan.Select(ctx, cr.GetDB(), &clients, sql)

	log.Println("GetAllClients is successful")

	return clients, nil
}
