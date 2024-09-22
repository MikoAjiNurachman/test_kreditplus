package LoanHandler

import (
	"kreditplus-api/handler"
	"kreditplus-api/model/dto"
	"kreditplus-api/service/LoanService"
	"net/http"
)

type loanHandler struct {
	service LoanService.LoanService
}

// HandleNoParam implements LoanHandler.
func (e *loanHandler) HandleNoParam(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handler.ServeTransactional(w, r, dto.Transactions{}, e.service.ApplyLoan)
	}
}

func NewLoanHandler(service LoanService.LoanService) LoanHandler {
	return &loanHandler{service: service}
}
