package CustomerService

import (
	"database/sql"
	"kreditplus-api/model/dao"
	"kreditplus-api/model/dto"
	"kreditplus-api/utils"
	"net/http"
	"time"
)

// UpdateData implements CustomerService.
func (e *customerService) UpdateData(r *http.Request, tx *sql.Tx, customer interface{}) (err dto.StandardError) {
	id := utils.ReadIDParam(r)

	pathImageKTP, err := e.GetImagePathCDN(r, "image_ktp")

	if err.Err != nil {
		return
	}

	pathImageSelfie, err := e.GetImagePathCDN(r, "image_selfie")

	if err.Err != nil {
		return
	}

	customerStruct, _ := customer.(*dto.Customer)

	err = e.validateUpdateRequest(customerStruct)
	if err.Err != nil {
		return
	}

	customerOnDB, err := e.customerRepo.DetailData(int64(id))

	if err.Err != nil {
		return
	}

	//optimistic locking with updated_at
	if customerOnDB.UpdatedAt.Time.Format(dto.DefaultDateTimeFormat) != customerStruct.UpdatedAtStr {
		err = dto.StandardError{}.GenerateLockDataError()
		return
	}

	customerRepo := dao.Customer{
		ID:          sql.NullInt64{Int64: int64(id)},
		NIK:         sql.NullString{String: customerStruct.NIK},
		FullName:    sql.NullString{String: customerStruct.FullName},
		LegalName:   sql.NullString{String: customerStruct.LegalName},
		BirthPlace:  sql.NullString{String: customerStruct.BirthPlace},
		BirthDate:   sql.NullTime{Time: customerStruct.BirthDate},
		Sallary:     sql.NullInt64{Int64: customerStruct.Sallary},
		ImageKtp:    sql.NullString{String: pathImageKTP},
		ImageSelfie: sql.NullString{String: pathImageSelfie},
		UpdatedAt:   sql.NullTime{Time: time.Now()},
		Deleted:     sql.NullBool{Bool: false},
	}

	err = e.customerRepo.UpdateData(tx, customerRepo)

	return
}

func (e customerService) validateUpdateRequest(customerStruct *dto.Customer) (err dto.StandardError) {
	return customerStruct.ValidateUpdate()
}
