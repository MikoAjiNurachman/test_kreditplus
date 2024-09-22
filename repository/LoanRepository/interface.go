package LoanRepository

import (
	"database/sql"
	"kreditplus-api/model/dao"
	"kreditplus-api/model/dto"
)

type LoanRepository interface {
	GetLimitTenor(customerID, tenor, year int64) (int64, dto.StandardError)
	SaveTransactions(*sql.Tx, dao.Transactions) dto.StandardError
}