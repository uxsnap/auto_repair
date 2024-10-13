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

type ClientsRepository struct {
	*BasePgRepository
}

func NewClientsRepo(client *db.Client) *ClientsRepository {
	return &ClientsRepository{
		NewBaseRepo(client, "clients"),
	}
}

func (cr *ClientsRepository) GetAll(ctx context.Context, params body.ClientBodyParams) ([]entity.ClientWithData, error) {
	log.Println(cr.Prefix + ": calling GetAll from repo")

	preSql := sq.Select("c.id, c.name, e.id, e.name, c.phone, c.has_documents, c.passport").
		From(cr.Prefix + " c").
		PlaceholderFormat(sq.Dollar).
		Join("employees e on e.id = employee_id").Where("c.is_deleted = false")

	if params.Name != "" {
		preSql = preSql.Where(sq.Like{"LOWER(c.name)": strings.ToLower("%" + params.Name + "%")})
	}

	if params.EmployeeName != "" {
		preSql = preSql.Where(sq.Like{"LOWER(e.name)": strings.ToLower("%" + params.EmployeeName + "%")})
	}

	if params.Passport != "" {
		preSql = preSql.Where(sq.Like{"LOWER(c.passport)": strings.ToLower("%" + params.Passport + "%")})
	}

	if params.Phone != "" {
		preSql = preSql.Where(sq.Like{"LOWER(c.phone)": strings.ToLower("%" + params.Phone + "%")})
	}

	sql, args, err := preSql.ToSql()

	if err != nil {
		log.Println(cr.Prefix + ": calling GetAll errored")
		return nil, err
	}

	clients := []entity.ClientWithData{}

	rows, rowsErr := cr.GetDB().Query(ctx, sql, args...)

	if rowsErr != nil {
		return nil, rowsErr
	}
	defer rows.Close()

	for rows.Next() {
		var client entity.ClientWithData

		err := rows.Scan(
			&client.Id,
			&client.Name,
			&client.Employee.Id,
			&client.Employee.Name,
			&client.Phone,
			&client.HasDocuments,
			&client.Passport,
		)

		if err != nil {
			return nil, err
		}

		clients = append(clients, client)
	}

	log.Println(cr.Prefix + ": returning from GetAll from repo")

	return clients, nil
}

func (cr *ClientsRepository) Create(ctx context.Context, client entity.Client) (uuid.UUID, error) {
	log.Println(cr.Prefix + ": calling Create from repo")

	sql, args, err := sq.
		Insert("clients").Columns(
		"id",
		"name",
		"employee_id",
		"phone",
		"has_documents",
		"passport",
		"is_deleted",
	).PlaceholderFormat(sq.Dollar).
		Values(client.Id, client.Name, client.EmployeeId, client.Phone, client.HasDocuments, client.Passport, false).
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

func (cr *ClientsRepository) Delete(ctx context.Context, clientID string) (uuid.UUID, error) {
	log.Println(cr.Prefix + ": calling Create from repo")

	sql, args, err := sq.
		Update("clients").
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

func (cr *ClientsRepository) Update(ctx context.Context, id uuid.UUID, clientData entity.Client) error {
	log.Println(cr.Prefix + ": calling Update from repo")

	sql, args, err := sq.
		Update("clients").
		SetMap(map[string]interface{}{
			"name":          clientData.Name,
			"employee_id":   clientData.EmployeeId,
			"phone":         clientData.Phone,
			"has_documents": clientData.HasDocuments,
			"passport":      clientData.Passport,
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
