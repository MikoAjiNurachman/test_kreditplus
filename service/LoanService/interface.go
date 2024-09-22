package LoanService

import (
	"database/sql"
	"kreditplus-api/model/dto"
	"net/http"
)

type LoanService interface {
	ApplyLoan(*http.Request,*sql.Tx, interface{}) (dto.StandardError) 
}