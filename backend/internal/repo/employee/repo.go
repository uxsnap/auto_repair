package repoEmployee

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/uxsnap/auto_repair/backend/internal/db"
	"github.com/uxsnap/auto_repair/backend/internal/entity"
	"github.com/uxsnap/auto_repair/backend/internal/repo"
)

type EmployeeRepository struct {
	*repo.BasePgRepository
}

func NewEmployeeRepo(client *db.Client) *EmployeeRepository {
	return &EmployeeRepository{
		repo.NewBaseRepo(client),
	}
}

func (cr *EmployeeRepository) GetAll(ctx context.Context) ([]entity.Employee, error) {
	log.Println("Employee: calling GetAll from repo")

	EmployeeSql := sq.Select("*").From("employee")

	sql, _, err := EmployeeSql.ToSql()

	if err != nil {
		log.Println("Employee: calling GetAll errored")
		return nil, err
	}

	employees := []entity.Employee{}

	pgxscan.Select(ctx, cr.GetDB(), &employees, sql)

	log.Println("Employee: returning from GetAll from repo")

	return employees, nil
}

func (cr *EmployeeRepository) Create(ctx context.Context, employee entity.Employee) error {
	log.Println("Employee: calling Create from repo")

	sql, args, err := sq.
		Insert("employee").Columns(
		"id",
		"name",
		"position",
		"employee_num",
	).PlaceholderFormat(sq.Dollar).
		Values(employee.Id, employee.Name, employee.Position, employee.EmployeeNum).
		ToSql()

	if err != nil {
		log.Println("Employee: calling Create errored")
		return err
	}

	if _, err = cr.GetDB().Exec(ctx, sql, args...); err != nil {
		log.Println("Employee: calling Create errored")
		return err
	}

	return nil
}

func (cr *EmployeeRepository) Update(ctx context.Context, clientData entity.Employee) error {
	return nil
}
