package LoginHandler

import "net/http"

type LoginHandler interface {
	Login(http.ResponseWriter, *http.Request)
}