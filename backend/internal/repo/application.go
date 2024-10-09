package repo

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
	"github.com/uxsnap/auto_repair/backend/internal/db"
	"github.com/uxsnap/auto_repair/backend/internal/entity"
)

type ApplicationsRepository struct {
	*BasePgRepository
}

func NewApplicationsRepo(client *db.Client) *ApplicationsRepository {
	return &ApplicationsRepository{
		NewBaseRepo(client, "applications"),
	}
}

func (cr *ApplicationsRepository) GetAll(ctx context.Context) ([]entity.Application, error) {
	log.Println(cr.Prefix + ": calling GetAll from repo")

	sql, _, err := sq.Select("id,employee_id,client_id,created_at,name,status,contract_id").
		From(cr.Prefix).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		log.Println(cr.Prefix + ": calling GetAll errored")
		return nil, err
	}

	applications := []entity.Application{}

	pgxscan.Select(ctx, cr.GetDB(), &applications, sql)

	log.Println(cr.Prefix + ": returning from GetAll from repo")

	return applications, nil
}

func (cr *ApplicationsRepository) Create(ctx context.Context, client entity.Application) (uuid.UUID, error) {
	log.Println(cr.Prefix + ": calling Create from repo")

	sql, args, err := sq.
		Insert("Applications").Columns(
		"id",
		"employee_id",
		"client_id", "created_at", "name", "status", "contract_id",
		"is_deleted",
	).PlaceholderFormat(sq.Dollar).
		Values(client.Id, client.EmployeeId, client.ClientId, client.CreatedAt, client.Name, client.Status, client.ContractId, false).
		ToSql()

	if err != nil {
		log.Println(cr.Prefix + ": calling Create errored")
		return uuid.Nil, err
	}

	if _, err = cr.GetDB().Exec(ctx, sql, args...); err != nil {
		log.Println(cr.Prefix + ": calling Create errored")
		return uuid.Nil, err
	}

	return client.Id.Bytes, nil
}

func (cr *ApplicationsRepository) Delete(ctx context.Context, clientID string) (uuid.UUID, error) {
	log.Println(cr.Prefix + ": calling Create from repo")

	sql, args, err := sq.
		Update(cr.Prefix).
		Set("is_deleted", true).
		Where(sq.Eq{"id": clientID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		log.Println(cr.Prefix + ": calling Create errored")
		return uuid.Nil, err
	}

	if _, err = cr.GetDB().Exec(ctx, sql, args...); err != nil {
		log.Println(cr.Prefix + ": calling Create errored")
		return uuid.Nil, err
	}

	return uuid.MustParse(clientID), nil
}
