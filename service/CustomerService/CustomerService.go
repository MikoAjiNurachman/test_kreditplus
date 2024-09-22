package CustomerService

import (
	"io"
	"kreditplus-api/model/dto"
	"kreditplus-api/repository/CustomerRepository"
	"net/http"
	"os"
	"path/filepath"
)

type customerService struct {
	customerRepo CustomerRepository.CustomerRepository
}

// GetImagePathCDN implements CustomerService.
func (e *customerService) GetImagePathCDN(r *http.Request, key string) (path string, err dto.StandardError) {

	if errS := r.ParseMultipartForm(10 << 20); errS != nil {
		err = dto.StandardError{
			Err:               errS,
			Code:              http.StatusBadRequest,
			AdditionalMessage: "invalid form-data",
		}
		return
	}
	fileContent, handler, errS := r.FormFile(key)

	if errS != nil {
		err = dto.StandardError{
			Err:               errS,
			Code:              http.StatusBadRequest,
			AdditionalMessage: "file not found",
		}
		return
	}

	defer func() {
		r.Body.Close()
		fileContent.Close()
	}()

	dir, _ := os.Getwd()

	path = filepath.Join(dir, "/assets", handler.Filename)

	targetFile, errS := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)

	if errS != nil {
		err = dto.StandardError{
			Err:               errS,
			Code:              http.StatusInternalServerError,
			AdditionalMessage: "failed to created target file",
		}
		return
	}

	defer targetFile.Close()

	if _, errS := io.Copy(targetFile, fileContent); errS != nil {
		err = dto.StandardError{
			Err:               errS,
			Code:              http.StatusInternalServerError,
			AdditionalMessage: "failed to write to target file",
		}
		return
	}

	return
}

func NewcustomerService(customerRepo CustomerRepository.CustomerRepository) CustomerService {
	return &customerService{customerRepo: customerRepo}
}
