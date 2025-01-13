package repo

import (
	"context"
	"log"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/uxsnap/auto_repair/backend/internal/body"
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

func (cr *ReceiptsRepository) GetAll(ctx context.Context, params body.ReceiptBodyParams) ([]entity.ReceiptWithData, error) {
	log.Println(cr.Prefix + ": calling GetAll from repo")

	preSql := sq.Select("r.id,r.created_at,c.id,c.name,r.sum").
		From(cr.Prefix + " r").
		PlaceholderFormat(sq.Dollar).
		Join("contracts c on c.id = r.contract_id")

	if params.ContractName != "" {
		preSql = preSql.Where(sq.Like{"LOWER(c.name)": strings.ToLower("%" + params.ContractName + "%")})
	}

	if params.MinSum != 0 {
		preSql = preSql.Where(sq.GtOrEq{"r.sum": params.MinSum})
	}

	if params.MaxSum != 0 {
		preSql = preSql.Where(sq.LtOrEq{"r.sum": params.MaxSum})
	}

	if params.MinCreatedAt != "" {
		preSql = preSql.Where(sq.GtOrEq{"r.created_at": params.MinCreatedAt})
	}

	if params.MaxCreatedAt != "" {
		preSql = preSql.Where(sq.LtOrEq{"r.created_at": params.MaxCreatedAt})
	}

	sql, args, err := preSql.ToSql()

	if err != nil {
		log.Println(cr.Prefix + ": calling GetAll errored")
		return nil, err
	}

	receipts := []entity.ReceiptWithData{}

	rows, rowsErr := cr.GetDB().Query(ctx, sql, args...)

	if rowsErr != nil {
		return nil, rowsErr
	}
	defer rows.Close()

	for rows.Next() {
		var receipt entity.ReceiptWithData

		err := rows.Scan(
			&receipt.Id,
			&receipt.CreatedAt,
			&receipt.Contract.Id,
			&receipt.Contract.Name,
			&receipt.Sum,
		)

		if err != nil {
			return nil, err
		}

		receipts = append(receipts, receipt)
	}

	log.Println(cr.Prefix + ": returning from GetAll from repo")

	return receipts, nil
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
