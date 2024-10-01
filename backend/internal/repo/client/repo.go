package repoClients

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	uuid "github.com/google/uuid"
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

func (cr *ClientsRepository) GetAll(ctx context.Context) ([]entity.Client, error) {
	log.Println("clients: calling GetAll from repo")

	clientsSql := sq.Select("*").From("clients")

	sql, _, err := clientsSql.ToSql()

	if err != nil {
		log.Println("clients: calling GetAll errored")
		return nil, err
	}

	clients := []entity.Client{}

	pgxscan.Select(ctx, cr.GetDB(), &clients, sql)

	log.Println("clients: returning from GetAll from repo")

	return clients, nil
}

func (cr *ClientsRepository) Create(ctx context.Context, client entity.Client) error {
	log.Println("clients: calling Create from repo")

	sql, args, err := sq.
		Insert("clients").Columns(
		"id",
		"name",
		"employeeId",
		"phone",
		"hasDocuments",
		"passport",
	).
		Values(uuid.New(), client.Name, client.EmployeeId, client.Phone, client.HasDocuments, client.Passport).
		ToSql()

	if err != nil {
		log.Println("clients: calling GetAll errored")
		return err
	}

	cr.GetDB().Query(ctx, sql, args...)

	log.Println("clients: calling GetAll errored")

	return nil
}

func (cr *ClientsRepository) Delete(ctx context.Context, clientID string) error {
	return nil
}

func (cr *ClientsRepository) Update(ctx context.Context, clientData entity.Client) error {
	return nil
}
