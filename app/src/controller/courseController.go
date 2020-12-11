package controller

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/felipesulzbach/exemplo-api-rest/app/src/controller/contract"
	"github.com/felipesulzbach/exemplo-api-rest/app/src/service"

)

func getAllCourse(w http.ResponseWriter, r *http.Request) {
	response, err := service.FindAllCourse()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		return
	}

	jsonOkResponse(w, response)
}

func getByIDCourse(w http.ResponseWriter, r *http.Request) {
	var contract contract.CourseContract
	id, err := contract.ValidatePath(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := service.FindByIDCourse(id)
	switch {
	case err == sql.ErrNoRows:
		var errorDesc bytes.Buffer
		errorDesc.WriteString("ERROR: No records found for id=")
		errorDesc.WriteString(strconv.FormatInt(id, 10))
		log.Println(errorDesc.String())
		json.NewEncoder(w).Encode(errorDesc.String())
		return
	case err != nil:
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		return
	default:
	}
	jsonOkResponse(w, response)
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	var contract contract.CourseContract
	entity, err := contract.ValidateBodyCreate(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := service.InsertCourse(entity)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		return
	}

	jsonCreatedResponse(w, response)
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	var contract contract.CourseContract
	entity, err := contract.ValidateBodyUpdate(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := service.UpdateCourse(entity)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		return
	}

	jsonOkResponse(w, response)
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	var contract contract.CourseContract
	id, err := contract.ValidatePath(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := service.DeleteCourse(id); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		return
	}

	jsonOkResponse(w, "")
}
