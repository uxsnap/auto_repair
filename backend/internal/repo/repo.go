package repo

import (
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
