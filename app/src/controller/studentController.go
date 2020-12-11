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

func getAllStudent(w http.ResponseWriter, r *http.Request) {
	response, err := service.FindAllStudent()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		return
	}

	jsonOkResponse(w, response)
}

func getByIDStudent(w http.ResponseWriter, r *http.Request) {
	var contract contract.StudentContract
	id, err := contract.ValidatePath(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := service.FindByIDStudent(id)
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

func createStudent(w http.ResponseWriter, r *http.Request) {
	var contract contract.StudentContract
	entity, err := contract.ValidateBodyCreate(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := service.InsertStudent(entity)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		return
	}

	jsonCreatedResponse(w, response)
}

func updateStudent(w http.ResponseWriter, r *http.Request) {
	var contract contract.StudentContract
	entity, err := contract.ValidateBodyUpdate(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := service.UpdateStudent(entity)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		return
	}

	jsonOkResponse(w, response)
}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	var contract contract.StudentContract
	id, err := contract.ValidatePath(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := service.DeleteStudent(id); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		return
	}

	jsonOkResponse(w, "")
}
