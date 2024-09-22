package LoanService

import (
	"database/sql"
	"kreditplus-api/model/dao"
	"kreditplus-api/model/dto"
	"kreditplus-api/repository/CustomerRepository"
	"kreditplus-api/repository/LimitCustomerRepository"
	"kreditplus-api/repository/LoanRepository"
	"math"
	"net/http"
	"time"
)

type loanService struct {
	loanRepo LoanRepository.LoanRepository
	customerRepo      CustomerRepository.CustomerRepository
	limitCustomerRepo LimitCustomerRepository.LimitCustomerRepository
}

// ApplyLoan implements LoanService.
func (e *loanService) ApplyLoan(r *http.Request, tx *sql.Tx, dataInterface interface{}) (err dto.StandardError)  {
	loanStruct := dataInterface.(*dto.Transactions)

	err = e.validateRequest(loanStruct)

	if err.Err!= nil {
        return
    }

	limit, err := e.loanRepo.GetLimitTenor(loanStruct.CustomerID, int64(loanStruct.Tenor), int64(loanStruct.Year))

	if loanStruct.CreditAmount > float64(limit) {
		err = dto.StandardError{}.GenerateInsufficientBalance()
        return
	}

	//TODO calculate cicilan dan bunga
	installment, rate, totalPayment := e.calcInstallmentAndRate(loanStruct)

	transactionDAO := dao.Transactions{
		CustomerID: sql.NullInt64{Int64: loanStruct.CustomerID},
		ContractNumber:    sql.NullString{String: loanStruct.ContractNumber},
        OtrPrice:          sql.NullFloat64{Float64: loanStruct.Otr},
        AdminFee:          sql.NullFloat64{Float64: loanStruct.AdminFee},
        InstallmentAmount: sql.NullFloat64{Float64: installment},
        RateAmount:        sql.NullFloat64{Float64: rate},
		AssetName:         sql.NullString{String: loanStruct.AssetName},
        TotalPayment:      sql.NullFloat64{Float64: totalPayment},
        CreatedAt:         sql.NullTime{Time: time.Now()},
        UpdatedAt:         sql.NullTime{Time: time.Now()},
        Deleted:           sql.NullBool{Bool: false},
	}

	err = e.loanRepo.SaveTransactions(tx, transactionDAO)

	if err.Err != nil {
		return
	}

	err = dto.StandardError{}.GenerateNoError()
	return
}

func (e *loanService) calcInstallmentAndRate(loanStruct *dto.Transactions) (installment, rate, totalPayment float64) {
		monthlyRate := float64(loanStruct.Rate) / 12.0 / 100.0

		// Hitung cicilan
		installment = float64(loanStruct.CreditAmount) * float64(monthlyRate) * math.Pow(1.0+float64(monthlyRate), float64(loanStruct.Tenor)) / (math.Pow(1.0+float64(monthlyRate), float64(loanStruct.Tenor)) - 1.0)
	
		// Hitung total bunga
		rate = (installment * float64(loanStruct.Tenor)) - float64(loanStruct.CreditAmount)

		// Hitung total pembayaran
		totalPayment = installment * float64(loanStruct.Tenor) + float64(loanStruct.AdminFee) + float64(loanStruct.Otr)//+

	return
}

func (e *loanService) validateRequest(loanStruct *dto.Transactions) dto.StandardError {
	return loanStruct.ValidateRequest()
}

func NewLoanService(loanRepository LoanRepository.LoanRepository, customerRepo CustomerRepository.CustomerRepository, limitCustomerRepo LimitCustomerRepository.LimitCustomerRepository) LoanService {
	return &loanService{
		loanRepo: loanRepository,
		customerRepo:      customerRepo,
		limitCustomerRepo: limitCustomerRepo,
	}
}
