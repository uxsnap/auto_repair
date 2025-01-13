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

type ApplicationsRepository struct {
	*BasePgRepository
}

func NewApplicationsRepo(client *db.Client) *ApplicationsRepository {
	return &ApplicationsRepository{
		NewBaseRepo(client, "applications"),
	}
}

func (cr *ApplicationsRepository) GetAll(ctx context.Context, params body.ApplicationBodyParams) ([]entity.ApplicationWithData, error) {
	log.Println(cr.Prefix + ": calling GetAll from repo")

	preSql := sq.Select("apps.id,e.id, e.name, c.id, c.name, apps.name, apps.created_at,apps.status,con.id, con.name").
		From(cr.Prefix + " apps").
		PlaceholderFormat(sq.Dollar).
		Join("employees e on e.id = apps.employee_id").
		Join("clients c on c.id = apps.client_id").
		Join("contracts con on con.id = apps.contract_id").
		Where("apps.is_deleted = false")

	if params.Name != "" {
		preSql = preSql.Where(sq.Like{"LOWER(apps.name)": strings.ToLower("%" + params.Name + "%")})
	}

	if params.EmployeeName != "" {
		preSql = preSql.Where(sq.Like{"LOWER(e.name)": strings.ToLower("%" + params.EmployeeName + "%")})
	}

	if params.ClientName != "" {
		preSql = preSql.Where(sq.Like{"LOWER(c.name)": strings.ToLower("%" + params.ClientName + "%")})
	}

	if params.ContractName != "" {
		preSql = preSql.Where(sq.Like{"LOWER(con.name)": strings.ToLower("%" + params.ContractName + "%")})
	}

	if params.Status != "" {
		preSql = preSql.Where(sq.Like{"LOWER(apps.status)": strings.ToLower("%" + params.Status + "%")})
	}

	if params.MinCreatedAt != "" {
		preSql = preSql.Where(sq.GtOrEq{"apps.created_at": params.MinCreatedAt})
	}

	if params.MaxCreatedAt != "" {
		preSql = preSql.Where(sq.LtOrEq{"apps.created_at": params.MaxCreatedAt})
	}

	sql, args, err := preSql.ToSql()

	if err != nil {
		log.Println(cr.Prefix + ": calling GetAll errored")
		return nil, err
	}

	applications := []entity.ApplicationWithData{}

	rows, rowsErr := cr.GetDB().Query(ctx, sql, args...)

	if rowsErr != nil {
		return nil, rowsErr
	}
	defer rows.Close()

	for rows.Next() {
		var application entity.ApplicationWithData

		err := rows.Scan(
			&application.Id,
			&application.Employee.Id,
			&application.Employee.Name,
			&application.Client.Id,
			&application.Client.Name,
			&application.Name,
			&application.CreatedAt,
			&application.Status,
			&application.Contract.Id,
			&application.Contract.Name,
		)

		if err != nil {
			return nil, err
		}

		applications = append(applications, application)
	}

	log.Println(cr.Prefix + ": returning from GetAll from repo")

	return applications, nil
}

func (cr *ApplicationsRepository) Create(ctx context.Context, client entity.Application) (uuid.UUID, error) {
	log.Println(cr.Prefix + ": calling Create from repo")

	sql, args, err := sq.
		Insert(cr.Prefix).Columns(
		"id",
		"employee_id",
		"client_id",
		"created_at",
		"name",
		"status",
		"contract_id",
		"is_deleted",
	).PlaceholderFormat(sq.Dollar).
		Values(client.Id, client.EmployeeId, client.ClientId, client.CreatedAt.Time, client.Name, client.Status, client.ContractId, false).
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

func (cr *ApplicationsRepository) Delete(ctx context.Context, clientID string) (uuid.UUID, error) {
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

func (cr *ApplicationsRepository) Update(ctx context.Context, id uuid.UUID, clientData entity.Application) error {
	log.Println(cr.Prefix + ": calling Update from repo")

	sql, args, err := sq.
		Update(cr.Prefix).
		SetMap(map[string]interface{}{
			"name":        clientData.Name,
			"employee_id": clientData.EmployeeId,
			"client_id":   clientData.ClientId,
			"contract_id": clientData.ContractId,
			"status":      clientData.Status,
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
