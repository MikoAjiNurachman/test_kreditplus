package CustomerService

import (
	"kreditplus-api/model/dto"
	"kreditplus-api/utils"
	"net/http"
)

func (e customerService) DeleteData(r *http.Request) (response dto.StandardResponse, err dto.StandardError) {
	id := utils.ReadIDParam(r)

	_, err = e.customerRepo.DetailData(int64(id))

	if err.Err != nil {
        return
    }

	err = e.customerRepo.DeleteData(int64(id))

	if err.Err != nil {
		return
	}

    err = dto.StandardError{}.GenerateNoError()
	return
}