package repo

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/uxsnap/auto_repair/backend/internal/db"
	"github.com/uxsnap/auto_repair/backend/internal/entity"
)

type ServicesRepository struct {
	*BasePgRepository
}

func NewServicesRepo(client *db.Client) *ServicesRepository {
	return &ServicesRepository{
		NewBaseRepo(client, "services"),
	}
}

func (cr *ServicesRepository) GetAll(ctx context.Context) ([]entity.Service, error) {
	log.Println(cr.Prefix + ": calling GetAll from repo")

	preSql := sq.Select("id,name").
		From(cr.Prefix).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := preSql.ToSql()

	if err != nil {
		log.Println(cr.Prefix + ": calling GetAll errored")
		return nil, err
	}

	services := []entity.Service{}

	rows, rowsErr := cr.GetDB().Query(ctx, sql, args...)

	if rowsErr != nil {
		return nil, rowsErr
	}
	defer rows.Close()

	for rows.Next() {
		var service entity.Service

		err := rows.Scan(
			&service.Id,
			&service.Name,
		)

		if err != nil {
			return nil, err
		}

		services = append(services, service)
	}

	log.Println(cr.Prefix + ": returning from GetAll from repo")

	return services, nil
}
