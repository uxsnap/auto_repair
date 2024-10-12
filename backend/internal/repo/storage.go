package repo

import (
	"context"
	"fmt"
	"log"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/uxsnap/auto_repair/backend/internal/body"
	"github.com/uxsnap/auto_repair/backend/internal/db"
	"github.com/uxsnap/auto_repair/backend/internal/entity"
)

type StoragesRepository struct {
	*BasePgRepository
}

func NewStoragesRepo(client *db.Client) *StoragesRepository {
	return &StoragesRepository{
		NewBaseRepo(client, "storages"),
	}
}

func (cr *StoragesRepository) GetAll(ctx context.Context, params body.StorageBodyParams) ([]entity.StorageWithData, error) {
	log.Println(cr.Prefix + ": calling GetAll from repo")

	preSql := sq.Select("s.id,s.storage_num,e.id,e.name,d.id,d.name,s.detail_count").
		From(cr.Prefix + " s").
		PlaceholderFormat(sq.Dollar).
		Join("details d on d.id = detail_id").
		Join("employees e on e.id = employee_id").
		Where("s.is_deleted = false")

	if params.StorageNum != "" {
		preSql = preSql.Where(sq.Like{"LOWER(storage_num)": strings.ToLower("%" + params.StorageNum + "%")})
	}

	if params.EmployeeName != "" {
		preSql = preSql.Where(sq.Like{"LOWER(e.name)": strings.ToLower("%" + params.EmployeeName + "%")})
	}

	if params.DetailName != "" {
		preSql = preSql.Where(sq.Like{"LOWER(d.name)": strings.ToLower("%" + params.DetailName + "%")})
	}

	sql, args, err := preSql.ToSql()

	fmt.Println(sql)

	if err != nil {
		log.Println(cr.Prefix + ": calling GetAll errored")
		return nil, err
	}

	Storages := []entity.StorageWithData{}

	rows, rowsErr := cr.GetDB().Query(ctx, sql, args...)

	if rowsErr != nil {
		return nil, rowsErr
	}
	defer rows.Close()

	for rows.Next() {
		var storage entity.StorageWithData

		err := rows.Scan(
			&storage.Id,
			&storage.StorageNum,
			&storage.Employee.Id,
			&storage.Employee.Name,
			&storage.Detail.Id,
			&storage.Detail.Name,
			&storage.DetailCount,
		)

		if err != nil {
			return nil, err
		}

		Storages = append(Storages, storage)
	}

	log.Println(cr.Prefix + ": returning from GetAll from repo")

	return Storages, nil
}

func (cr *StoragesRepository) Create(ctx context.Context, client entity.Storage) (uuid.UUID, error) {
	log.Println(cr.Prefix + ": calling Create from repo")

	sql, args, err := sq.
		Insert(cr.Prefix).Columns(
		"id", "employee_id", "storage_num", "detail_id", "detail_count", "is_deleted",
	).PlaceholderFormat(sq.Dollar).
		Values(client.Id, client.EmployeeId, client.StorageNum, client.DetailId, client.DetailCount, false).
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

func (cr *StoragesRepository) Delete(ctx context.Context, clientID string) (uuid.UUID, error) {
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

func (cr *StoragesRepository) Update(ctx context.Context, id uuid.UUID, clientData body.CreateStorageBody) error {
	log.Println(cr.Prefix + ": calling Update from repo")

	sql, args, err := sq.
		Update(cr.Prefix).
		SetMap(map[string]interface{}{
			"storage_num":  clientData.StorageNum,
			"employee_id":  clientData.EmployeeId,
			"detail_id":    clientData.DetailId,
			"detail_count": clientData.DetailCount,
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

	log.Printf(cr.Prefix+": updated %d rows", res.RowsAffected())

	return nil
}
