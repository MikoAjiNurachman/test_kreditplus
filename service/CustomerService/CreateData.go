package CustomerService

import (
	"database/sql"
	"kreditplus-api/model/dao"
	"kreditplus-api/model/dto"
	"net/http"
	"time"
)

func (e customerService) UpsertData(r *http.Request, tx *sql.Tx, customer interface{}) (err dto.StandardError) {

	customerStruct, _ := customer.(*dto.Customer)

	pathImageKTP, err := e.GetImagePathCDN(r, "image_ktp")

	if err.Err != nil {
		return
	}

	pathImageSelfie, err := e.GetImagePathCDN(r, "image_selfie")

	if err.Err != nil {
		return
	}

	err = e.validateInsertRequest(customerStruct)

	if err.Err != nil {
		return
	}

	

	customerRepo := dao.Customer{
		NIK: sql.NullString{String: customerStruct.NIK},
		FullName: sql.NullString{String: customerStruct.FullName},
		LegalName: sql.NullString{String: customerStruct.LegalName},
		BirthPlace: sql.NullString{String: customerStruct.BirthPlace},
		BirthDate: sql.NullTime{Time: customerStruct.BirthDate},
		Sallary: sql.NullInt64{Int64: customerStruct.Sallary},
		ImageKtp: sql.NullString{String: pathImageKTP},
		ImageSelfie: sql.NullString{String: pathImageSelfie},
		CreatedAt: sql.NullTime{Time: time.Now()},
		UpdatedAt: sql.NullTime{Time: time.Now()},
        Deleted: sql.NullBool{Bool: false},
	}

	err = e.customerRepo.UpsertData(tx, customerRepo)

	return
}

func (e customerService) validateInsertRequest(customerStruct *dto.Customer) (err dto.StandardError) {
	return customerStruct.ValidateInsert()
}
