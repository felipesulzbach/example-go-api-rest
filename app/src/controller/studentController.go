package controller

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/felipesulzbach/exemplo-api-rest/app/src/model"
	"github.com/felipesulzbach/exemplo-api-rest/app/src/service"
	"github.com/gorilla/mux"

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
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
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
		log.Panic(err)
		return
	default:
	}

	jsonOkResponse(w, response)
}

func insertStudent(w http.ResponseWriter, r *http.Request) {
	var entity model.Student
	_ = json.NewDecoder(r.Body).Decode(&entity)

	id, err := service.InsertStudent(entity)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		return
	}

	jsonCreatedResponse(w, id)
}

func updateStudent(w http.ResponseWriter, r *http.Request) {
	var entity model.Student
	_ = json.NewDecoder(r.Body).Decode(&entity)

	if err := service.UpdateStudent(entity); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		return
	}
}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	// TODO
}
