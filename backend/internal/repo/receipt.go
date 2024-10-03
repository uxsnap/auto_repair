package repo

import (
	"context"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
	"github.com/uxsnap/auto_repair/backend/internal/db"
	"github.com/uxsnap/auto_repair/backend/internal/entity"
)

type ReceiptsRepository struct {
	*BasePgRepository
}

func NewReceiptsRepo(client *db.Client) *ReceiptsRepository {
	return &ReceiptsRepository{
		NewBaseRepo(client, "receipts"),
	}
}

func (cr *ReceiptsRepository) GetAll(ctx context.Context) ([]entity.Receipt, error) {
	log.Println(cr.Prefix + ": calling GetAll from repo")

	sql, _, err := sq.Select("id,created_at,contract_id,sum").
		From(cr.Prefix).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		log.Println(cr.Prefix + ": calling GetAll errored")
		return nil, err
	}

	var Receipts []entity.Receipt

	pgxscan.Select(ctx, cr.GetDB(), &Receipts, sql)

	log.Println(cr.Prefix + ": returning from GetAll from repo")

	return Receipts, nil
}

func (cr *ReceiptsRepository) Create(ctx context.Context, client entity.Receipt) (uuid.UUID, error) {
	log.Println(cr.Prefix + ": calling Create from repo")

	sql, args, err := sq.
		Insert(cr.Prefix).Columns(
		"id",
		"created_at",
		"contract_id",
		"sum",
		"is_deleted",
	).PlaceholderFormat(sq.Dollar).
		Values(client.Id, time.Now(), client.ContractId, client.Sum, false).
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

func (cr *ReceiptsRepository) Update(ctx context.Context, id uuid.UUID, clientData entity.Receipt) error {
	log.Println(cr.Prefix + ": calling Update from repo")

	sql, args, err := sq.
		Update(cr.Prefix).
		SetMap(map[string]interface{}{
			"sum": clientData.Sum,
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

	log.Printf("Receipts: updated %d rows", res.RowsAffected())

	return nil
}

func (cr *ReceiptsRepository) Delete(ctx context.Context, id string) (uuid.UUID, error) {
	log.Println(cr.Prefix + ": calling Create from repo")

	sql, args, err := sq.
		Update(cr.Prefix).
		Set("is_deleted", true).
		Where(sq.Eq{"id": id}).
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

	return uuid.MustParse(id), nil
}
