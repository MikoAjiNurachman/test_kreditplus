package container

import (
	"fmt"
	"kreditplus-api/config"
	"kreditplus-api/handler/CustomerHandler"
	"kreditplus-api/handler/LimitCustomerHandler"
	"kreditplus-api/handler/LoanHandler"
	"kreditplus-api/handler/LoginHandler"
	"kreditplus-api/handler/MiddlewareHandler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRouter(container *Container) {

	r := mux.NewRouter()

	loginHandler := LoginHandler.NewLoginHandler(container.LoginService)
	r.HandleFunc("/login", loginHandler.Login).Methods(http.MethodPost, http.MethodOptions)

	protectedR := r.PathPrefix("/api").Subrouter()
	protectedR.Use(MiddlewareHandler.TokenVerifyMiddleware)

	customerHandler := CustomerHandler.NewCustomerHandler(container.CustomerService)
	protectedR.HandleFunc("/customer", customerHandler.HandleNoParam).Methods(http.MethodGet, http.MethodPost, http.MethodOptions)
	protectedR.HandleFunc("/customer/{ID}", customerHandler.HandleWithParam).Methods(http.MethodGet, http.MethodPut, http.MethodDelete, http.MethodOptions)

	limitCustomerHandler := LimitCustomerHandler.NewLimitCustomerHandler(container.LimitCustomerService)
	protectedR.HandleFunc("/limitcustomer", limitCustomerHandler.HandleNoParam).Methods(http.MethodPost, http.MethodGet, http.MethodOptions)

	loanHandler := LoanHandler.NewLoanHandler(container.LoanService)
	protectedR.HandleFunc("/applyloan", loanHandler.HandleNoParam).Methods(http.MethodPost, http.MethodOptions)


	log.Printf(`server running on port %v`, config.ServerConfig.ServerPort)

	http.ListenAndServe(fmt.Sprintf(`:%s`, config.ServerConfig.ServerPort), r)
}
