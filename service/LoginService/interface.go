package LoginService

import (
	"kreditplus-api/model/dto"
	"net/http"
)

type LoginService interface {
	LoginService(*http.Request) (dto.StandardResponse, dto.StandardError)
}