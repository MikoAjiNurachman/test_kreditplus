package LoanHandler

import "net/http"

type LoanHandler interface {
	HandleNoParam(w http.ResponseWriter, r *http.Request)
}
