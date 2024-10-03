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

type EmployeesRepository struct {
	*BasePgRepository
}

func NewEmployeesRepo(client *db.Client) *EmployeesRepository {
	return &EmployeesRepository{
		NewBaseRepo(client, "employees"),
	}
}

func (cr *EmployeesRepository) GetAll(ctx context.Context) ([]entity.Employee, error) {
	log.Println(cr.Prefix + ": calling GetAll from repo")

	EmployeeSql := sq.Select("*").From("employees")

	sql, _, err := EmployeeSql.ToSql()

	if err != nil {
		log.Println(cr.Prefix + ": calling GetAll errored")
		return nil, err
	}

	employees := []entity.Employee{}

	pgxscan.Select(ctx, cr.GetDB(), &employees, sql)

	log.Println(cr.Prefix + ": returning from GetAll from repo")

	return employees, nil
}

func (cr *EmployeesRepository) Create(ctx context.Context, employee entity.Employee) (uuid.UUID, error) {
	log.Println(cr.Prefix + ": calling Create from repo")

	sql, args, err := sq.
		Insert("employees").Columns(
		"id",
		"name",
		"position",
		"employee_num",
	).PlaceholderFormat(sq.Dollar).
		Values(employee.Id, employee.Name, employee.Position, employee.EmployeeNum).
		ToSql()

	if err != nil {
		log.Println(cr.Prefix + ": calling Create errored")
		return uuid.Nil, err
	}

	if _, err = cr.GetDB().Exec(ctx, sql, args...); err != nil {
		log.Println(cr.Prefix + ": calling Create errored")
		return uuid.Nil, err
	}

	return employee.Id.Bytes, nil
}

func (cr *EmployeesRepository) Update(ctx context.Context, id uuid.UUID, clientData entity.Employee) error {
	log.Println(cr.Prefix + ": calling Update from repo")

	sql, args, err := sq.
		Update("employees").
		SetMap(map[string]interface{}{
			"name":         clientData.Name,
			"position":     clientData.Position,
			"employee_num": clientData.EmployeeNum,
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

	log.Printf("clients: updated %d rows", res.RowsAffected())

	return nil
}
