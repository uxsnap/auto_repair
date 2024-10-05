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

type StoragesRepository struct {
	*BasePgRepository
}

func NewStoragesRepo(client *db.Client) *StoragesRepository {
	return &StoragesRepository{
		NewBaseRepo(client, "storages"),
	}
}

func (cr *StoragesRepository) GetAll(ctx context.Context) ([]entity.Storage, error) {
	log.Println(cr.Prefix + ": calling GetAll from repo")

	sql, _, err := sq.Select("id,employee_id,storage_num").
		From(cr.Prefix).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		log.Println(cr.Prefix + ": calling GetAll errored")
		return nil, err
	}

	var Storages []entity.Storage

	pgxscan.Select(ctx, cr.GetDB(), &Storages, sql)

	log.Println(cr.Prefix + ": returning from GetAll from repo")

	return Storages, nil
}

func (cr *StoragesRepository) Create(ctx context.Context, client entity.Storage) (uuid.UUID, error) {
	log.Println(cr.Prefix + ": calling Create from repo")

	sql, args, err := sq.
		Insert("Storages").Columns(
		"id", "employee_id", "storage_num",
	).PlaceholderFormat(sq.Dollar).
		Values(client.Id, client.EmployeeId, client.StorageNum).
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
