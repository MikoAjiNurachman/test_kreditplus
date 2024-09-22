package LimitCustomerService

import (
	"database/sql"
	"kreditplus-api/model/dao"
	"kreditplus-api/model/dto"
	"kreditplus-api/repository/CustomerRepository"
	"kreditplus-api/repository/LimitCustomerRepository"
	"net/http"
	"time"
)

type limitCustomerService struct {
	repo         LimitCustomerRepository.LimitCustomerRepository
	customerRepo CustomerRepository.CustomerRepository
}

// ListLimitCustomer implements LimitCustomerService.
func (e *limitCustomerService) ListLimitCustomer(*http.Request) (dto.StandardResponse, dto.StandardError) {
	result, err := e.repo.ListLimitCustomer()
	if err.Err != nil {
		return dto.StandardResponse{}, err
	}

	content := e.convertListRepoToContent(result)

	return dto.StandardResponse{
		Content: content,
	}, dto.StandardError{}.GenerateNoError()
}

func (e *limitCustomerService) convertListRepoToContent(listData []dao.ListLimitCustomer) (result []dto.ListLimitCustomer) {

	for _, v := range listData {
		temp := dto.ListLimitCustomer{
			CustomerNIK: v.CustomerNIK.String,
			CustomerName: v.CustomerName.String,
			ImageSelfie: v.ImageSelfie.String,
			Year: v.Year.Int64,
            Tenor1: v.Tenor1.Int64,
            Tenor2: v.Tenor2.Int64,
            Tenor3: v.Tenor3.Int64,
            Tenor4: v.Tenor4.Int64,
			Tenor5: v.Tenor5.Int64,
			Tenor6: v.Tenor6.Int64,
			Tenor7: v.Tenor7.Int64,
			Tenor8: v.Tenor8.Int64,
			Tenor9: v.Tenor9.Int64,
			Tenor10: v.Tenor10.Int64,
			Tenor11: v.Tenor11.Int64,
			Tenor12: v.Tenor12.Int64,
		}
		result = append(result, temp)
	}

	return
}

// UpsertData implements LimitCustomerService.
func (e *limitCustomerService) UpsertData(r *http.Request, tx *sql.Tx, dataInterface interface{}) (err dto.StandardError) {
	limitCustomer := dataInterface.(*dto.LimitCustomer)

	e.validateUpsertRequest(limitCustomer)

	//validate customer on db
	_, err = e.customerRepo.DetailData(limitCustomer.CustomerID)

	if err.Err != nil {
		return
	}

	limitCustomerDao := dao.LimitCustomer{
		CustomerID: sql.NullInt64{Int64: limitCustomer.CustomerID},
		Year:       sql.NullInt64{Int64: limitCustomer.Year},
		Tenor1:     sql.NullInt64{Int64: limitCustomer.Tenor1},
		Tenor2:     sql.NullInt64{Int64: limitCustomer.Tenor2},
		Tenor3:     sql.NullInt64{Int64: limitCustomer.Tenor3},
		Tenor4:     sql.NullInt64{Int64: limitCustomer.Tenor4},
		Tenor5:     sql.NullInt64{Int64: limitCustomer.Tenor5},
		Tenor6:     sql.NullInt64{Int64: limitCustomer.Tenor6},
		Tenor7:     sql.NullInt64{Int64: limitCustomer.Tenor7},
		Tenor8:     sql.NullInt64{Int64: limitCustomer.Tenor8},
		Tenor9:     sql.NullInt64{Int64: limitCustomer.Tenor9},
		Tenor10:    sql.NullInt64{Int64: limitCustomer.Tenor10},
		Tenor11:    sql.NullInt64{Int64: limitCustomer.Tenor11},
		Tenor12:    sql.NullInt64{Int64: limitCustomer.Tenor12},
		CreatedAt:  sql.NullTime{Time: time.Now()},
		UpdatedAt:  sql.NullTime{Time: time.Now()},
		Deleted:    sql.NullBool{Bool: false},
	}

	err = e.repo.UpsertData(tx, limitCustomerDao)

	if err.Err != nil {
		return
	}

	err = dto.StandardError{}.GenerateNoError()
	return
}

func (e *limitCustomerService) validateUpsertRequest(limitCustomer *dto.LimitCustomer) dto.StandardError {
	return limitCustomer.ValidateUpsert()
}

func NewLimitCustomerService(repo LimitCustomerRepository.LimitCustomerRepository, customerRepo CustomerRepository.CustomerRepository) LimitCustomerService {
	return &limitCustomerService{
		repo:         repo,
		customerRepo: customerRepo,
	}
}
