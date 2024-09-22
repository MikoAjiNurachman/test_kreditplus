package LimitCustomerHandler

import (
	"kreditplus-api/handler"
	"kreditplus-api/model/dto"
	"kreditplus-api/service/LimitCustomerService"
	"net/http"
)

type limitCustomerHandler struct {
	service LimitCustomerService.LimitCustomerService
}

func (e limitCustomerHandler) HandleNoParam(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handler.ServeTransactional(w, r, dto.LimitCustomer{}, e.service.UpsertData)
	case http.MethodGet:
		handler.Serve(w, r, e.service.ListLimitCustomer)
	}
}

func NewLimitCustomerHandler(service LimitCustomerService.LimitCustomerService) LimitCustomerHandler {
	return &limitCustomerHandler{service: service}
}
