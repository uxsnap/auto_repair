package repoClients

import (
	"context"
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
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

	sql, _, err := sq.Select("id, employee_id, name, phone, passport, has_documents, is_deleted").
		From("clients").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		log.Println("clients: calling GetAll errored")
		return nil, err
	}

	var clients []entity.Client

	pgxscan.Select(ctx, cr.GetDB(), &clients, sql)

	log.Println("clients: returning from GetAll from repo")

	fmt.Println(clients)

	return clients, nil
}

func (cr *ClientsRepository) Create(ctx context.Context, client entity.Client) (uuid.UUID, error) {
	log.Println("clients: calling Create from repo")

	sql, args, err := sq.
		Insert("clients").Columns(
		"id",
		"name",
		"employee_id",
		"phone",
		"has_documents",
		"passport",
		"is_deleted",
	).PlaceholderFormat(sq.Dollar).
		Values(client.Id, client.Name, client.EmployeeId, client.Phone, client.HasDocuments, client.Passport, false).
		ToSql()

	if err != nil {
		log.Println("clients: calling Create errored")
		return uuid.Nil, err
	}

	if _, err = cr.GetDB().Exec(ctx, sql, args...); err != nil {
		log.Println("clients: calling Create errored")
		return uuid.Nil, err
	}

	return client.Id.Bytes, nil
}

func (cr *ClientsRepository) Delete(ctx context.Context, clientID string) (uuid.UUID, error) {
	log.Println("clients: calling Create from repo")

	sql, args, err := sq.
		Update("clients").
		Set("is_deleted", true).
		Where(sq.Eq{"id": clientID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		log.Println("clients: calling Create errored")
		return uuid.Nil, err
	}

	if _, err = cr.GetDB().Exec(ctx, sql, args...); err != nil {
		log.Println("clients: calling Create errored")
		return uuid.Nil, err
	}

	return uuid.MustParse(clientID), nil
}

func (cr *ClientsRepository) Update(ctx context.Context, id uuid.UUID, clientData entity.Client) error {
	log.Println("clients: calling Update from repo")

	sql, args, err := sq.
		Update("clients").
		SetMap(map[string]interface{}{
			"name":          clientData.Name,
			"employee_id":   clientData.EmployeeId,
			"phone":         clientData.Phone,
			"has_documents": clientData.HasDocuments,
			"passport":      clientData.Passport,
		}).
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		log.Println("clients: calling Update errored")
		return err
	}

	res, err := cr.GetDB().Exec(ctx, sql, args...)

	if err != nil {
		log.Println("clients: calling Update errored")
		return err
	}

	log.Printf("clients: updated %d rows", res.RowsAffected())

	return nil
}
