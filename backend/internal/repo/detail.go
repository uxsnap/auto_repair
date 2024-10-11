package repo

import (
	"context"
	"fmt"
	"log"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
	"github.com/uxsnap/auto_repair/backend/internal/body"
	"github.com/uxsnap/auto_repair/backend/internal/db"
	"github.com/uxsnap/auto_repair/backend/internal/entity"
)

type DetailsRepository struct {
	*BasePgRepository
}

func NewDetailsRepo(client *db.Client) *DetailsRepository {
	return &DetailsRepository{
		NewBaseRepo(client, "details"),
	}
}

func (cr *DetailsRepository) GetAll(ctx context.Context, params body.DetailBodyParams) ([]entity.Detail, error) {
	log.Println(cr.Prefix + ": calling GetAll from repo")

	preSql := sq.Select("id, name, price, type").
		From(cr.Prefix).
		PlaceholderFormat(sq.Dollar).
		Where("is_deleted = false")

	if params.Name != "" {
		preSql = preSql.Where(sq.Like{"LOWER(name)": strings.ToLower("%" + params.Name + "%")})
	}

	if params.MinPrice != 0 {
		preSql = preSql.Where(sq.GtOrEq{"price": params.MinPrice})
	}

	if params.MaxPrice != 0 {
		preSql = preSql.Where(sq.LtOrEq{"price": params.MaxPrice})
	}

	if params.Type != "" {
		preSql = preSql.Where(sq.Eq{"type": params.Type})
	}

	sql, args, err := preSql.ToSql()

	if err != nil {
		log.Println(cr.Prefix + ": calling GetAll errored")
		return nil, err
	}

	details := []entity.Detail{}

	pgxscan.Select(ctx, cr.GetDB(), &details, sql, args...)

	fmt.Println(sql, args)

	log.Println(cr.Prefix + ": returning from GetAll from repo")

	return details, nil
}

func (cr *DetailsRepository) Create(ctx context.Context, client entity.Detail) (uuid.UUID, error) {
	log.Println(cr.Prefix + ": calling Create from repo")

	sql, args, err := sq.
		Insert(cr.Prefix).Columns(
		"id",
		"name",
		"type",
		"price",
		"is_deleted",
	).PlaceholderFormat(sq.Dollar).
		Values(client.Id, client.Name, client.Type, client.Price, false).
		ToSql()

	fmt.Println(sql)

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

func (cr *DetailsRepository) Delete(ctx context.Context, clientID string) (uuid.UUID, error) {
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
