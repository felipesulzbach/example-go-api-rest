package service

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/felipesulzbach/exemplo-api-rest/app/src/model"
	"github.com/felipesulzbach/exemplo-api-rest/app/src/repository"
	"github.com/gorilla/mux"

)

// FindAllStudent ...
func FindAllStudent(w http.ResponseWriter, r *http.Request) {
	list, err := repository.FindAllStudent()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		return
	}

	for _, item := range list {
		log.Println(item.ToString())
	}

	jsonOkResponse(w, list)
}

// FindByIDStudent ...
func FindByIDStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		return
	}

	entity, err := repository.FindByIDStudent(id)
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

	log.Println(entity.ToString())
	jsonOkResponse(w, entity)
}

// InsertStudent ...
func InsertStudent(w http.ResponseWriter, r *http.Request) {
	var entity model.Student
	_ = json.NewDecoder(r.Body).Decode(&entity)

	id, err := repository.InsertStudent(entity)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		return
	}
	jsonCreatedResponse(w, id)
}

// UpdateStudent ...
func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	var entity model.Student
	_ = json.NewDecoder(r.Body).Decode(&entity)

	if err := repository.UpdateStudent(entity); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		return
	}
}

// DeleteStudent ...
func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	// TODO
}
