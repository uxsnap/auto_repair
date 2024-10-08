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

type VehiclesRepository struct {
	*BasePgRepository
}

func NewVehiclesRepo(client *db.Client) *VehiclesRepository {
	return &VehiclesRepository{
		NewBaseRepo(client, "vehicles"),
	}
}

func (cr *VehiclesRepository) GetAll(ctx context.Context) ([]entity.Vehicle, error) {
	log.Println(cr.Prefix + ": calling GetAll from repo")

	sql, _, err := sq.Select("id, client_id, vehicle_number, brand, model, is_deleted").
		From(cr.Prefix).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		log.Println(cr.Prefix + ": calling GetAll errored")
		return nil, err
	}

	var Vehicles []entity.Vehicle

	pgxscan.Select(ctx, cr.GetDB(), &Vehicles, sql)

	log.Println(cr.Prefix + ": returning from GetAll from repo")

	return Vehicles, nil
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
	).PlaceholderFormat(sq.Dollar).
		Values(client.Id, client.ClientId, client.Brand, client.Model, client.VehicleNumber).
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
