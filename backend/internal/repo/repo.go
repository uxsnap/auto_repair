package repo

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/uxsnap/auto_repair/backend/internal/db"
)

type BasePgRepository struct {
	DB *db.Client
}

func NewBaseRepo(dbc *db.Client) *BasePgRepository {
	return &BasePgRepository{
		DB: dbc,
	}
}

func (bpr *BasePgRepository) GetDB() *pgxpool.Pool {
	return bpr.DB.GetDb()
}
