package dto

import (
	"errors"
	"net/http"
	"regexp"
	"time"
)

type Customer struct {
	ID           int64     `json:"id"`
	NIK          string    `json:"nik"`
	FullName     string    `json:"full_name"`
	LegalName    string    `json:"legal_name"`
	BirthPlace   string    `json:"birth_place"`
	BirthDate    time.Time `json:"-"`
	BirthDateStr string    `json:"birth_date"`
	Sallary      int64     `json:"sallary"`
	ImageKtp     string    `json:"image_ktp"`
	ImageSelfie  string    `json:"image_selfie"`
	CreatedAtStr string    `json:"created_at"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAtStr string    `json:"updated_at"`
	UpdatedAt    time.Time `json:"-"`
	Deleted      bool      `json:"deleted"`
}

type ListCustomer struct {
	NIK       string `json:"nik"`
	FullName  string `json:"full_name"`
	LegalName string `json:"legal_name"`
	Deleted   bool   `json:"deleted"`
}

func (e *Customer) ValidateInsert() (err StandardError) {
	err = e.validateNIK()

	if err.Err != nil {
		return
	}

	err = e.validateFullName()
	if err.Err != nil {
		return
	}
	err = e.validateLegalName()
	if err.Err != nil {
		return
	}
	err = e.validateBirthPlace()
	if err.Err != nil {
		return
	}
	err = e.validateBirthDate()
	if err.Err != nil {
		return
	}

	err = e.validateSallary()
	if err.Err != nil {
		return
	}
	err = StandardError{}.GenerateNoError()
	return
}

func (e *Customer) validateNIK() (err StandardError) {

	if e.NIK == "" {
		err = StandardError{}.GenerateEmptyField("NIK")
		return
	}

	rgx := regexp.MustCompile(regexNIK)

	if !rgx.MatchString(e.NIK) {
		err = StandardError{
			Err:               errors.New("NIK must be number and at least 16 characters"),
			Code:              http.StatusBadRequest,
			AdditionalMessage: "NIK must be number and at least 16 characters",
		}
		return
	}

	err = StandardError{}.GenerateNoError()
	return
}

func (e *Customer) validateFullName() (err StandardError) {

	if e.FullName == "" {
		err = StandardError{}.GenerateEmptyField("Full Name")
		return
	}
	err = StandardError{}.GenerateNoError()
	return
}

func (e *Customer) validateLegalName() (err StandardError) {

	if e.LegalName == "" {
		err = StandardError{}.GenerateEmptyField("Legal Name")
		return
	}
	err = StandardError{}.GenerateNoError()
	return
}

func (e *Customer) validateBirthPlace() (err StandardError) {

	if e.BirthPlace == "" {
		err = StandardError{}.GenerateEmptyField("Birth Place")
		return
	}
	err = StandardError{}.GenerateNoError()
	return
}

func (e *Customer) validateBirthDate() (err StandardError) {
	var errS error

	if e.BirthDateStr == "" {
		err = StandardError{}.GenerateEmptyField("Birth Date")
		return
	}

	e.BirthDate, errS = time.Parse(DefaultDateFormat, e.BirthDateStr)

	if errS != nil {
		err = StandardError{
			Err:               errS,
			Code:              http.StatusBadRequest,
			AdditionalMessage: "Invalid Birth Date Format",
		}
		return
	}

	err = StandardError{}.GenerateNoError()
	return
}

func (e *Customer) validateSallary() (err StandardError) {

	if e.Sallary < 1 {
		err = StandardError{
			Err:               errors.New("Sallary must be greater than 0"),
			Code:              http.StatusBadRequest,
			AdditionalMessage: "Sallary must be greater than 0",
		}
		return
	}

	err = StandardError{}.GenerateNoError()
	return
}

func (e *Customer) ValidateUpdate() (err StandardError) {

	err = e.validateFullName()
	if err.Err != nil {
		return
	}

	err = e.validateLegalName()
	if err.Err != nil {
		return
	}

	err = e.validateBirthPlace()
	if err.Err != nil {
		return
	}

	err = e.validateBirthDate()
	if err.Err != nil {
		return
	}

	err = e.validateSallary()
	if err.Err != nil {
		return
	}

	if e.UpdatedAtStr == "" {
		err = StandardError{}.GenerateEmptyField("updated_at")
		return
	}

	e.UpdatedAt, _ = time.Parse(DefaultDateTimeFormat, e.UpdatedAtStr)

	err = StandardError{}.GenerateNoError()
	return
}
