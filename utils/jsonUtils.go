package utils

import (
	"encoding/json"
	"kreditplus-api/model/dto"
	"net/http"
)

func ReadBody(r *http.Request, content interface{}) (err dto.StandardError) {
	errS := json.NewDecoder(r.Body).Decode(content)
	// defer r.Body.Close()
	if errS != nil {
		err = dto.StandardError{
			Err:               errS,
			Code:              http.StatusInternalServerError,
			AdditionalMessage: "invalid decode json",
		}
		return
	}
	return
}
func ReadBodyMultipart(r *http.Request, content interface{}) (err dto.StandardError) {

	value := r.FormValue("form_content")

	errS := json.Unmarshal([]byte(value), content)

	if errS != nil {
		err = dto.StandardError{
			Err:               errS,
			Code:              http.StatusInternalServerError,
			AdditionalMessage: "invalid decode json",
		}
		return
	}
	return
}