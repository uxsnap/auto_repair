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

type VehiclesRepository struct {
	*BasePgRepository
}

func NewVehiclesRepo(client *db.Client) *VehiclesRepository {
	return &VehiclesRepository{
		NewBaseRepo(client, "vehicles"),
	}
}

func (cr *VehiclesRepository) GetAll(ctx context.Context, params body.VehicleBodyParams) ([]entity.VehicleWithData, error) {
	log.Println(cr.Prefix + ": calling GetAll from repo")

	preSql := sq.Select("v.id, c.id, c.name, v.vehicle_number, v.brand, v.model").
		From(cr.Prefix + " v").
		PlaceholderFormat(sq.Dollar).
		Join("clients c on c.id = v.client_id").
		Where("v.is_deleted = false")

	if params.VehicleNumber != "" {
		preSql = preSql.Where(sq.Like{"LOWER(v.vehicle_number)": strings.ToLower("%" + params.VehicleNumber + "%")})
	}

	if params.Brand != "" {
		preSql = preSql.Where(sq.Like{"LOWER(v.brand)": strings.ToLower("%" + params.Brand + "%")})
	}

	if params.Model != "" {
		preSql = preSql.Where(sq.Like{"LOWER(v.model)": strings.ToLower("%" + params.Model + "%")})
	}

	sql, args, err := preSql.ToSql()

	if err != nil {
		log.Println(cr.Prefix + ": calling GetAll errored")
		return nil, err
	}

	vehicles := []entity.VehicleWithData{}

	rows, rowsErr := cr.GetDB().Query(ctx, sql, args...)

	if rowsErr != nil {
		return nil, rowsErr
	}
	defer rows.Close()

	for rows.Next() {
		var vehicle entity.VehicleWithData

		err := rows.Scan(
			&vehicle.Id,
			&vehicle.Client.Id,
			&vehicle.Client.Name,
			&vehicle.VehicleNumber,
			&vehicle.Brand,
			&vehicle.Model,
		)

		if err != nil {
			return nil, err
		}

		vehicles = append(vehicles, vehicle)
	}

	log.Println(cr.Prefix + ": returning from GetAll from repo")

	return vehicles, nil
}

func (cr *VehiclesRepository) Create(ctx context.Context, client entity.Vehicle) (uuid.UUID, error) {
	log.Println(cr.Prefix + ": calling Create from repo")

	sql, args, err := sq.
		Insert(cr.Prefix).Columns(
		"id",
		"client_id",
		"brand",
		"model",
		"vehicle_number",
		"is_deleted",
	).PlaceholderFormat(sq.Dollar).
		Values(client.Id, client.ClientId, client.Brand, client.Model, client.VehicleNumber, false).
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

func (cr *VehiclesRepository) Delete(ctx context.Context, clientID string) (uuid.UUID, error) {
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
