package handler

import (
	"database/sql"
	"encoding/json"
	"kreditplus-api/config"
	"kreditplus-api/model/dto"
	"kreditplus-api/utils"
	"log"
	"net/http"
	"reflect"
)

func ServeTransactional(w http.ResponseWriter, r *http.Request, bodyType interface{},
	serviceFunc func(*http.Request, *sql.Tx, interface{}) dto.StandardError) {
	var (
		response dto.StandardResponse
		db       = config.ServerConfig.DBCon
	)
	body := reflect.New(reflect.TypeOf(bodyType)).Interface()
	err := utils.ReadBody(r, body)
	tx, errS := db.Begin()

	_, errS = db.Exec("SET search_path TO sample_schema, public")
	if errS != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	if errS != nil {
		log.Fatalf(`error create transactional %v`, errS)
	}
	err = serviceFunc(r, tx, body)
	defer func() {
		if err.Err != nil {
			errS = tx.Rollback()
			if errS != nil {
				response.Code = err.Code
				response.Message = err.AdditionalMessage
				writeResponse(w, response)
				return
			}
			response.Code = err.Code
			response.Message = err.AdditionalMessage
			writeResponse(w, response)
			return
		}
		errS = tx.Commit()
		if errS != nil {
			response.Code = err.Code
			response.Message = err.AdditionalMessage
			writeResponse(w, response)
			return
		}
		response.Code = http.StatusOK
		writeResponse(w, response)
	}()
}

func ServeTransactionalWithMultipart(w http.ResponseWriter, r *http.Request, bodyType interface{},
	serviceFunc func(*http.Request, *sql.Tx, interface{}) dto.StandardError) {
	var (
		response dto.StandardResponse
		db       = config.ServerConfig.DBCon
	)
	body := reflect.New(reflect.TypeOf(bodyType)).Interface()
	err := utils.ReadBodyMultipart(r, body)

	tx, errS := db.Begin()

	_, errS = db.Exec("SET search_path TO sample_schema, public")
	if errS != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	if errS != nil {
		log.Fatalf(`error create transactional %v`, errS)
	}
	err = serviceFunc(r, tx, body)
	defer func() {
		if err.Err != nil {
			errS = tx.Rollback()
			if errS != nil {
				response.Code = err.Code
				response.Message = err.AdditionalMessage
				writeResponse(w, response)
				return
			}
			response.Code = err.Code
			response.Message = err.AdditionalMessage
			writeResponse(w, response)
			return
		}
		errS = tx.Commit()
		if errS != nil {
			response.Code = err.Code
			response.Message = err.AdditionalMessage
			writeResponse(w, response)
			return
		}
		response.Code = http.StatusOK
		writeResponse(w, response)
	}()
}

func Serve(w http.ResponseWriter, r *http.Request, serviceFunc func(*http.Request) (dto.StandardResponse, dto.StandardError)) {
	response, err := serviceFunc(r)
	if err.Err != nil {
		response.Code = err.Code
		response.Message = err.AdditionalMessage
		writeResponse(w, response)
		return
	}
	response.Code = http.StatusOK
	response.Message = "Success"
	writeResponse(w, response)
}

func writeResponse(w http.ResponseWriter, content dto.StandardResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(content.Code)
	err := json.NewEncoder(w).Encode(content)
	if err != nil {
		log.Fatalf(`error marshal response`)
	}
}
