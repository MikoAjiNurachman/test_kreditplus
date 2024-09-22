package LimitCustomerService

import (
	"database/sql"
	"kreditplus-api/model/dto"
	"net/http"
)

type LimitCustomerService interface {
	UpsertData(*http.Request, *sql.Tx, interface{}) dto.StandardError
	ListLimitCustomer(*http.Request) (dto.StandardResponse, dto.StandardError)
}