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

func getAllTeacher(w http.ResponseWriter, r *http.Request) {
	response, err := service.FindAllTeacher()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		return
	}

	jsonOkResponse(w, response)
}

func getByIDTeacher(w http.ResponseWriter, r *http.Request) {
	var contract contract.TeacherContract
	id, err := contract.ValidatePath(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := service.FindByIDTeacher(id)
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

func createTeacher(w http.ResponseWriter, r *http.Request) {
	var contract contract.TeacherContract
	entity, err := contract.ValidateBodyCreate(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := service.InsertTeacher(entity)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		return
	}

	jsonCreatedResponse(w, response)
}

func updateTeacher(w http.ResponseWriter, r *http.Request) {
	var contract contract.TeacherContract
	entity, err := contract.ValidateBodyUpdate(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := service.UpdateTeacher(entity)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		return
	}

	jsonOkResponse(w, response)
}

func deleteTeacher(w http.ResponseWriter, r *http.Request) {
	var contract contract.TeacherContract
	id, err := contract.ValidatePath(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := service.DeleteTeacher(id); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		return
	}

	jsonOkResponse(w, "")
}
