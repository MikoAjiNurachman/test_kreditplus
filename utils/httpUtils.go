package utils

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ReadIDParam(r *http.Request) int {
	idStr := mux.Vars(r)["ID"]
	id, _ := strconv.Atoi(idStr)
	return id
}
