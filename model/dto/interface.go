package dto

import (
	"errors"
	"fmt"
	"net/http"
)

const regexNIK = `^\d{16}$`

const DefaultDateFormat = "2006-01-02"
const DefaultDateTimeFormat = "2006-01-02T15:04:05Z"

type StandardResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Content interface{} `json:"content"`
}

type StandardError struct {
	Err               error
	Code              int
	AdditionalMessage string
}

func (e StandardError) GenerateLockDataError() StandardError {
	return StandardError{
		Err:               errors.New("cant process data, already use by other"),
		Code:              http.StatusLocked,
		AdditionalMessage: "Can't Process Data, Already use by Other",
	}
}

func (e StandardError) GenerateContainsCharacter(max int) StandardError {
	return StandardError{
		Err:               errors.New(fmt.Sprintf("field must contains %d characters", max)),
		Code:              http.StatusBadRequest,
		AdditionalMessage: fmt.Sprintf("Field must contains %d characters", max),
	}
}

func (e StandardError) GenerateGreaterLessThan(min, max int) StandardError {
	return StandardError{
		Err:               errors.New(fmt.Sprintf("field must be greater than %d and less than %d", min, max)),
		Code:              http.StatusBadRequest,
		AdditionalMessage: fmt.Sprintf("Field must be greater than %d and less than %d", min, max),
	}
}

func (e StandardError) GenerateEmptyField(field string) StandardError {
	return StandardError{
		Err:               errors.New(fmt.Sprintf("empty field %s", field)),
		Code:              http.StatusBadRequest,
		AdditionalMessage: fmt.Sprintf("Empty Field %s", field),
	}
}

func (e StandardError) GenerateInvalidRequest() StandardError {
	return StandardError{
		Err:               errors.New("invalid request"),
		Code:              400,
		AdditionalMessage: "Invalid Request",
	}
}

func (e StandardError) GenerateInsufficientBalance() StandardError {
	return StandardError{
        Err:               errors.New("insufficient balance"),
        Code:              403,
        AdditionalMessage: "Insufficient Balance",
    }
}

func (e StandardError) GenerateInvalidCredentials() StandardError {
	return StandardError{
        Err:               errors.New("invalid credentials"),
        Code:              http.StatusUnauthorized,
        AdditionalMessage: "Invalid Credentials",
    }
}

func (e StandardError) GenerateInternalServerError(errS error) StandardError {
	return StandardError{
		Err:               errS,
		Code:              http.StatusInternalServerError,
		AdditionalMessage: errS.Error(),
	}
}

func (e StandardError) GenerateNoError() StandardError {
	return StandardError{
		Code:              http.StatusOK,
		AdditionalMessage: "Success",
	}
}

func (e StandardError) GenerateNotFound(field string) StandardError {
	return StandardError{
		Err:               errors.New(field),
		Code:              http.StatusBadRequest,
		AdditionalMessage: field,
	}
}
