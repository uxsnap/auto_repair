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

type ActsRepository struct {
	*BasePgRepository
}

func NewActsRepo(client *db.Client) *ActsRepository {
	return &ActsRepository{
		NewBaseRepo(client, "acts"),
	}
}

func (cr *ActsRepository) GetAll(ctx context.Context) ([]entity.Act, error) {
	log.Println(cr.Prefix + ": calling GetAll from repo")

	sql, _, err := sq.Select("id,name,application_id,created_at,service_id").
		From(cr.Prefix).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		log.Println(cr.Prefix + ": calling GetAll errored")
		return nil, err
	}

	var Acts []entity.Act

	pgxscan.Select(ctx, cr.GetDB(), &Acts, sql)

	log.Println(cr.Prefix + ": returning from GetAll from repo")

	return Acts, nil
}

func (cr *ActsRepository) Create(ctx context.Context, client entity.Act) (uuid.UUID, error) {
	log.Println(cr.Prefix + ": calling Create from repo")

	sql, args, err := sq.
		Insert("Acts").Columns(
		"id", "name", "application_id", "created_at", "service_id", "is_deleted",
	).PlaceholderFormat(sq.Dollar).
		Values(client.Id, client.Name, client.ApplicationId, client.CreatedAt, client.ServiceId, false).
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

func (cr *ActsRepository) Delete(ctx context.Context, clientID string) (uuid.UUID, error) {
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
