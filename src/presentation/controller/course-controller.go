package controller

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/felipesulzbach/example-go-api-rest/src/domain/service"
	contract "github.com/felipesulzbach/example-go-api-rest/src/infra/contract/contractcourse"
	"github.com/felipesulzbach/example-go-api-rest/src/presentation/helper"

)

// GetAllCourse ...
func GetAllCourse(w http.ResponseWriter, r *http.Request) {
	response, err := service.FindAllCourse()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		return
	}

	helper.HTTPResponseOK(w, response)
}

// GetByIDCourse ...
func GetByIDCourse(w http.ResponseWriter, r *http.Request) {
	requestPath, err := contract.GetPath(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	isValid, errValidation := helper.ValidateRequestContract(requestPath)
	if errValidation != nil {
		http.Error(w, fmt.Sprintf("%s", errValidation), http.StatusBadRequest)
		return
	}
	log.Println("Validate:", isValid)

	response, err := service.FindByIDCourse(requestPath)
	switch {
	case err == sql.ErrNoRows:
		var errorDesc bytes.Buffer
		errorDesc.WriteString("ERROR: No records found for id=")
		errorDesc.WriteString(strconv.FormatInt(requestPath.ID, 10))
		log.Println(errorDesc.String())
		json.NewEncoder(w).Encode(errorDesc.String())
		return
	case err != nil:
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		return
	default:
	}

	helper.HTTPResponseOK(w, response)
}

// CreateCourse ...
func CreateCourse(w http.ResponseWriter, r *http.Request) {
	requestBody, err := contract.GetCreateRequestBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	isValid, errValidation := helper.ValidateRequestContract(requestBody)
	if errValidation != nil {
		http.Error(w, fmt.Sprintf("%s", errValidation), http.StatusBadRequest)
		return
	}
	log.Println("Validate:", isValid)

	response, err := service.InsertCourse(requestBody)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		return
	}

	helper.HTTPResponseCreate(w, response)
}

// UpdateCourse ...
func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	requestBody, err := contract.GetUpdateRequestBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	isValid, errValidation := helper.ValidateRequestContract(requestBody)
	if errValidation != nil {
		http.Error(w, fmt.Sprintf("%s", errValidation), http.StatusBadRequest)
		return
	}
	log.Println("Validate:", isValid)

	response, err := service.UpdateCourse(requestBody)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		return
	}

	helper.HTTPResponseOK(w, response)
}

// DeleteCourse ...
func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	requestPath, err := contract.GetPath(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	isValid, errValidation := helper.ValidateRequestContract(requestPath)
	if errValidation != nil {
		http.Error(w, fmt.Sprintf("%s", errValidation), http.StatusBadRequest)
		return
	}
	log.Println("Validate:", isValid)

	if err := service.DeleteCourse(requestPath); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		return
	}

	helper.HTTPResponseOK(w, "")
}
