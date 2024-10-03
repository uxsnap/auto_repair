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

type ContractsRepository struct {
	*BasePgRepository
}

func NewContractsRepo(client *db.Client) *ContractsRepository {
	return &ContractsRepository{
		NewBaseRepo(client, "contracts"),
	}
}

func (cr *ContractsRepository) GetAll(ctx context.Context) ([]entity.Contract, error) {
	log.Println(cr.Prefix + ": calling GetAll from repo")

	sql, _, err := sq.Select("id,name,sum,created_at,signed_at,status_id").
		From(cr.Prefix).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		log.Println(cr.Prefix + ": calling GetAll errored")
		return nil, err
	}

	var contracts []entity.Contract

	pgxscan.Select(ctx, cr.GetDB(), &contracts, sql)

	log.Println(cr.Prefix + ": returning from GetAll from repo")

	return contracts, nil
}

func (cr *ContractsRepository) Create(ctx context.Context, client entity.Contract) (uuid.UUID, error) {
	log.Println(cr.Prefix + ": calling Create from repo")

	sql, args, err := sq.
		Insert(cr.Prefix).Columns(
		"id",
		"name",
		"sum",
		"created_at",
		"signed_at",
		"status_id",
	).PlaceholderFormat(sq.Dollar).
		Values(client.Id, client.Name, client.Sum, client.CreatedAt, nil, client.StatusId).
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

func (cr *ContractsRepository) Update(ctx context.Context, id uuid.UUID, clientData entity.Contract) error {
	log.Println(cr.Prefix + ": calling Update from repo")

	sql, args, err := sq.
		Update(cr.Prefix).
		SetMap(map[string]interface{}{
			"name":      clientData.Name,
			"sum":       clientData.Sum,
			"signed_at": clientData.SignedAt,
			"status_id": clientData.StatusId,
		}).
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		log.Println(cr.Prefix + ": calling Update errored")
		return err
	}

	res, err := cr.GetDB().Exec(ctx, sql, args...)

	if err != nil {
		log.Println(cr.Prefix + ": calling Update errored")
		return err
	}

	log.Printf("Contracts: updated %d rows", res.RowsAffected())

	return nil
}
