package LimitCustomerRepository

import (
	"database/sql"
	"kreditplus-api/model/dao"
	"kreditplus-api/model/dto"
)

type LimitCustomerRepository interface {
	UpsertData(*sql.Tx, dao.LimitCustomer) dto.StandardError
	ListLimitCustomer() ([]dao.ListLimitCustomer, dto.StandardError)
}