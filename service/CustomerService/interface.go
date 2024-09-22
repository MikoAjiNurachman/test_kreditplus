package CustomerService

import (
	"database/sql"
	"kreditplus-api/model/dto"
	"net/http"
)

type CustomerService interface {
	UpsertData(*http.Request, *sql.Tx, interface{}) dto.StandardError
	GetImagePathCDN(*http.Request, string) (string, dto.StandardError)
	ListData(*http.Request) (dto.StandardResponse, dto.StandardError)
	DetailData(*http.Request) (dto.StandardResponse, dto.StandardError)
	UpdateData(*http.Request, *sql.Tx, interface{}) dto.StandardError
	DeleteData(*http.Request) (dto.StandardResponse, dto.StandardError)
}