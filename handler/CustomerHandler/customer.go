package CustomerHandler

import (
	"kreditplus-api/handler"
	"kreditplus-api/model/dto"
	"kreditplus-api/service/CustomerService"
	"net/http"
)

type customerHandler struct {

	service CustomerService.CustomerService
}

func (e customerHandler) HandleWithParam(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handler.Serve(w, r, e.service.DetailData)
	case http.MethodPut:
		handler.ServeTransactionalWithMultipart(w, r, dto.Customer{}, e.service.UpdateData)
	case http.MethodDelete:
		handler.Serve(w, r, e.service.DeleteData)
	}
}

func (e customerHandler) HandleNoParam(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handler.ServeTransactionalWithMultipart(w, r, dto.Customer{}, e.service.UpsertData)
	case http.MethodGet:
		handler.Serve(w, r, e.service.ListData)
	}
}

func NewCustomerHandler(service CustomerService.CustomerService) CustomerHandler {
	return &customerHandler{service: service}
}
