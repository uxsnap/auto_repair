package entity

import (
	"github.com/jackc/pgtype"
)

type Storage struct {
	Id          pgtype.UUID `json:"id"`
	EmployeeId  pgtype.UUID `json:"employeeId"`
	StorageNum  string      `json:"storageNum"`
	DetailId    pgtype.UUID `json:"detailId"`
	DetailCount int         `json:"detailCount"`
	IsDeleted   bool        `json:"isDeleted"`
}

type StorageWithData struct {
	Id          pgtype.UUID `json:"id"`
	Employee    IdWithName  `json:"employee"`
	Detail      IdWithName  `json:"detail"`
	StorageNum  string      `json:"storageNum"`
	DetailCount int         `json:"detailCount"`
	IsDeleted   bool        `json:"isDeleted"`
}
