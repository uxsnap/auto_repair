package repo

import (
	"context"
	"log"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/uxsnap/auto_repair/backend/internal/body"
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

func (cr *ActsRepository) GetAll(ctx context.Context, params body.ActBodyParams) ([]entity.ActWithData, error) {
	log.Println(cr.Prefix + ": calling GetAll from repo")

	preSql := sq.Select("act.id,act.name,a.id,a.name,act.created_at,s.id,s.name").
		From(cr.Prefix + " act").
		PlaceholderFormat(sq.Dollar).
		Join("applications a on a.id = act.application_id").
		Join("services s on s.id = act.service_id").
		Where("act.is_deleted = false")

	if params.Name != "" {
		preSql = preSql.Where(sq.Like{"LOWER(c.name)": strings.ToLower("%" + params.Name + "%")})
	}

	if params.ApplicationName != "" {
		preSql = preSql.Where(sq.Like{"LOWER(a.name)": strings.ToLower("%" + params.ApplicationName + "%")})
	}

	if params.ServiceName != "" {
		preSql = preSql.Where(sq.Like{"LOWER(s.name)": strings.ToLower("%" + params.ServiceName + "%")})
	}

	if params.MinCreatedAt != "" {
		preSql = preSql.Where(sq.GtOrEq{"act.created_at": params.MinCreatedAt})
	}

	if params.MaxCreatedAt != "" {
		preSql = preSql.Where(sq.LtOrEq{"act.created_at": params.MaxCreatedAt})
	}

	sql, args, err := preSql.ToSql()

	if err != nil {
		log.Println(cr.Prefix + ": calling GetAll errored")
		return nil, err
	}

	acts := []entity.ActWithData{}

	rows, rowsErr := cr.GetDB().Query(ctx, sql, args...)

	if rowsErr != nil {
		return nil, rowsErr
	}
	defer rows.Close()

	for rows.Next() {
		var client entity.ActWithData

		err := rows.Scan(
			&client.Id,
			&client.Name,
			&client.Application.Id,
			&client.Application.Name,
			&client.CreatedAt,
			&client.Service.Id,
			&client.Service.Name,
		)

		if err != nil {
			return nil, err
		}

		acts = append(acts, client)
	}

	log.Println(cr.Prefix + ": returning from GetAll from repo")

	return acts, nil
}

func (cr *ActsRepository) Create(ctx context.Context, client entity.Act) (uuid.UUID, error) {
	log.Println(cr.Prefix + ": calling Create from repo")

	sql, args, err := sq.
		Insert(cr.Prefix).Columns(
		"id", "name", "application_id", "created_at", "service_id", "is_deleted",
	).PlaceholderFormat(sq.Dollar).
		Values(client.Id, client.Name, client.ApplicationId, client.CreatedAt.Time, client.ServiceId, false).
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
