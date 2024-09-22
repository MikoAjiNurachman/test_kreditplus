package LimitCustomerHandler

import "net/http"

type LimitCustomerHandler interface {
	HandleNoParam(w http.ResponseWriter, r *http.Request)
}
