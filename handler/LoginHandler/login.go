package LoginHandler

import (
	"kreditplus-api/handler"
	"kreditplus-api/service/LoginService"
	"net/http"
)

type loginHandler struct {
	service LoginService.LoginService
}

// Login implements LoginHandler.
func (e *loginHandler) Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodPost:
            handler.Serve(w, r, e.service.LoginService)
        default:
            w.WriteHeader(http.StatusMethodNotAllowed)
            return
	}
}

func NewLoginHandler(service LoginService.LoginService) LoginHandler {
	return &loginHandler{service: service}
}
