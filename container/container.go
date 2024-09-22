package container

import (
	"database/sql"
	"kreditplus-api/config"
	"kreditplus-api/repository/CustomerRepository"
	"kreditplus-api/repository/LimitCustomerRepository"
	"kreditplus-api/repository/LoanRepository"
	"kreditplus-api/service/CustomerService"
	"kreditplus-api/service/LimitCustomerService"
	"kreditplus-api/service/LoanService"
	"kreditplus-api/service/LoginService"
)

type Container struct {
	DB              *sql.DB
	CustomerRepo    CustomerRepository.CustomerRepository
	CustomerService CustomerService.CustomerService
	LimitCustomerRepo LimitCustomerRepository.LimitCustomerRepository
	LimitCustomerService LimitCustomerService.LimitCustomerService
	LoanService LoanService.LoanService
	LoanRepo LoanRepository.LoanRepository
	LoginService LoginService.LoginService
}

func InitContainer() *Container {
	// Initialize repositories
	customerRepository := CustomerRepository.NewCustomerRepository(config.ServerConfig.DBCon)
	limitCustomer := LimitCustomerRepository.NewLimitCustomerRepository(config.ServerConfig.DBCon)
	loanRepository := LoanRepository.NewLoanRepository(config.ServerConfig.DBCon)

	// Initialize services
	customerService := CustomerService.NewcustomerService(customerRepository)
	limitCustomerService := LimitCustomerService.NewLimitCustomerService(limitCustomer, customerRepository)
	loanService := LoanService.NewLoanService(loanRepository, customerRepository, limitCustomer)

	loginService := LoginService.NewLoginService()

	return &Container{
		DB:              config.ServerConfig.DBCon,
		CustomerRepo:    customerRepository,
		LimitCustomerRepo: limitCustomer,
		CustomerService: customerService,
		LimitCustomerService: limitCustomerService,
		LoanService: loanService,
		LoanRepo: loanRepository,
		LoginService: loginService,
	}
}
