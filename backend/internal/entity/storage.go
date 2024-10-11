package entity

import (
	"github.com/jackc/pgtype"
)

type Storage struct {
	Id         pgtype.UUID `json:"id"`
	EmployeeId pgtype.UUID `json:"employeeId"`
	StorageNum string      `json:"storageNum"`
	DetailId   pgtype.UUID `json:"detailId"`
	IsDeleted  bool        `json:"isDeleted"`
}
