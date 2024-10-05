package entity

import (
	"github.com/jackc/pgtype"
)

type Storage struct {
	Id         pgtype.UUID `json:"id"`
	EmployeeId pgtype.UUID `json:"employee_id"`
	StorageNum string      `json:"storage_num"`
	IsDeleted  bool        `json:"is_deleted"`
}
