package CustomerService

import (
	"kreditplus-api/model/dao"
	"kreditplus-api/model/dto"
	"kreditplus-api/utils"
	"net/http"
)

// DetailData implements CustomerService.
func (e *customerService) DetailData(r *http.Request) (dto.StandardResponse, dto.StandardError) {
	id := utils.ReadIDParam(r)

	result, err := e.customerRepo.DetailData(int64(id))

	if err.Err != nil {
        return dto.StandardResponse{}, err
    }

	content := e.convertRepoToContentDetail(result)

	return dto.StandardResponse{
        Content: content,
    }, dto.StandardError{}.GenerateNoError()
}

func (e *customerService) convertRepoToContentDetail(repo dao.Customer) (result dto.Customer) {
	result.ID = repo.ID.Int64
	result.NIK = repo.NIK.String
	result.FullName = repo.FullName.String
	result.LegalName = repo.LegalName.String
	result.BirthDate = repo.BirthDate.Time
	result.BirthDateStr = repo.BirthDate.Time.Format(dto.DefaultDateFormat)
	result.BirthPlace = repo.BirthPlace.String
	result.Sallary = repo.Sallary.Int64
	result.ImageKtp = repo.ImageKtp.String
	result.ImageSelfie = repo.ImageSelfie.String
	result.CreatedAtStr = repo.CreatedAt.Time.Format(dto.DefaultDateTimeFormat)
	result.UpdatedAtStr = repo.UpdatedAt.Time.Format(dto.DefaultDateTimeFormat)
	result.Deleted = repo.Deleted.Bool
	return
}