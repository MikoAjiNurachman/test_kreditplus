package CustomerService

import (
	"kreditplus-api/model/dao"
	"kreditplus-api/model/dto"
	"net/http"
)

// ListData implements CustomerService.
func (e *customerService) ListData(*http.Request) (dto.StandardResponse, dto.StandardError) {
	result, err := e.customerRepo.ListData()
	if err.Err != nil {
		return dto.StandardResponse{}, err
	}

	content := e.convertListRepoToContent(result)

	return dto.StandardResponse{
		Content: content,
	}, dto.StandardError{}.GenerateNoError()
}

func (e *customerService) convertListRepoToContent(listData []dao.Customer) (result []dto.ListCustomer) {

	for _, v := range listData {
		var temp dto.ListCustomer
		temp.NIK = v.NIK.String
		temp.FullName = v.FullName.String
		temp.LegalName = v.LegalName.String
		temp.Deleted = v.Deleted.Bool
		result = append(result, temp)
	}

	return
}
