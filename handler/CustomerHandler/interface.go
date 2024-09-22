package CustomerHandler

import "net/http"

type CustomerHandler interface {
	HandleWithParam(w http.ResponseWriter, r *http.Request)
	HandleNoParam(w http.ResponseWriter, r *http.Request)
}
