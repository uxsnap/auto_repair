package repo

import "github.com/uxsnap/auto_repair/backend/internal/db"

type BasePgRepository struct {
	dbc *db.Client
}

func NewBaseRepo(dbc *db.Client) *BasePgRepository {
	return &BasePgRepository{
		dbc: dbc,
	}
}
