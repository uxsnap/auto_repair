package repo

import (
	"context"
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/uxsnap/auto_repair/backend/internal/db"
)

type BasePgRepository struct {
	DB     *db.Client
	Prefix string
}

func NewBaseRepo(dbc *db.Client, prefix string) *BasePgRepository {
	return &BasePgRepository{
		DB:     dbc,
		Prefix: prefix,
	}
}

func (bpr *BasePgRepository) GetDB() *pgxpool.Pool {
	return bpr.DB.GetDb()
}

type BaseEntity interface {
	Table() string
	Columns() string
	// IdColumnName() string
	// Scan(row pgx.Row) error
	// Values() []interface{}
	// ColumnsForUpdate() []string
	// ValuesForUpdate() []interface{}
}

func (bpr *BasePgRepository) GetAll(ctx context.Context, baseEntity BaseEntity, dst interface{}) error {
	log.Println(bpr.Prefix + ": calling GetAll from repo")

	sql, args, err := sq.Select(baseEntity.Columns()).
		From(baseEntity.Table()).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		log.Println(bpr.Prefix + ": calling GetAll errored")
		return err
	}

	pgxscan.Select(ctx, bpr.GetDB(), &dst, sql, args...)

	fmt.Println(dst)

	log.Println(bpr.Prefix + ": returning from GetAll from repo")

	return nil
}
