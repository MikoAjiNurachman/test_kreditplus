package CustomerRepository

import (
	"database/sql"
	"kreditplus-api/model/dao"
	"kreditplus-api/model/dto"
)

type CustomerRepository interface {
	UpsertData(*sql.Tx, dao.Customer) dto.StandardError
	ListData() ([]dao.Customer, dto.StandardError)
	DetailData(id int64) (dao.Customer, dto.StandardError)
	UpdateData(*sql.Tx, dao.Customer) dto.StandardError
	DeleteData(id int64) dto.StandardError
}