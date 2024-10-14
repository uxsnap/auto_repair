package repo

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/uxsnap/auto_repair/backend/internal/body"
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

func (cr *ContractsRepository) GetAll(ctx context.Context, params body.ContractBodyParams) ([]entity.Contract, error) {
	log.Println(cr.Prefix + ": calling GetAll from repo")

	preSql := sq.Select("id,name,sum,created_at,signed_at,status").
		From(cr.Prefix).
		PlaceholderFormat(sq.Dollar).
		Where("is_deleted = false")

	if params.Name != "" {
		preSql = preSql.Where(sq.Like{"LOWER(name)": strings.ToLower("%" + params.Name + "%")})
	}

	if params.Status != "" {
		preSql = preSql.Where(sq.Like{"LOWER(status)": strings.ToLower("%" + params.Status + "%")})
	}

	if params.MinSum != 0 {
		preSql = preSql.Where(sq.GtOrEq{"sum": params.MinSum})
	}

	if params.MaxSum != 0 {
		preSql = preSql.Where(sq.LtOrEq{"sum": params.MaxSum})
	}

	sql, args, err := preSql.ToSql()

	if err != nil {
		log.Println(cr.Prefix + ": calling GetAll errored")
		return nil, err
	}

	contracts := []entity.Contract{}

	rows, rowsErr := cr.GetDB().Query(ctx, sql, args...)

	if rowsErr != nil {
		return nil, rowsErr
	}
	defer rows.Close()

	for rows.Next() {
		var contract entity.Contract

		//id,name,sum,created_at,signed_at,status

		err := rows.Scan(
			&contract.Id,
			&contract.Name,
			&contract.Sum,
			&contract.CreatedAt,
			&contract.SignedAt,
			&contract.Status,
		)

		if err != nil {
			return nil, err
		}

		contracts = append(contracts, contract)
	}

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
		"status",
		"is_deleted",
	).PlaceholderFormat(sq.Dollar).
		Values(client.Id, client.Name, client.Sum, time.Now(), nil, client.Status, false).
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

func (cr *ContractsRepository) Delete(ctx context.Context, clientID string) (uuid.UUID, error) {
	log.Println(cr.Prefix + ": calling Delete from repo")

	sql, args, err := sq.
		Update(cr.Prefix).
		Set("is_deleted", true).
		Where(sq.Eq{"id": clientID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	fmt.Println(sql)

	if err != nil {
		log.Println(cr.Prefix + ": calling Delete errored")
		return uuid.Nil, err
	}

	if _, err = cr.GetDB().Exec(ctx, sql, args...); err != nil {
		log.Println(cr.Prefix + ": calling Delete errored")
		return uuid.Nil, err
	}

	return uuid.MustParse(clientID), nil
}

func (cr *ContractsRepository) Update(ctx context.Context, id uuid.UUID, clientData entity.Contract) error {
	log.Println(cr.Prefix + ": calling Update from repo")

	sql, args, err := sq.
		Update(cr.Prefix).
		SetMap(map[string]interface{}{
			"name":      clientData.Name,
			"sum":       clientData.Sum,
			"signed_at": clientData.SignedAt.Time,
			"status":    clientData.Status,
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
