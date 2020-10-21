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

// FindAllSchoolClass ...
func FindAllSchoolClass(w http.ResponseWriter, r *http.Request) {
	list, err := repository.FindAllSchoolClass()
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

// FindByIDSchoolClass ...
func FindByIDSchoolClass(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		return
	}

	entity, err := repository.FindByIDSchoolClass(id)
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

// InsertSchoolClass ...
func InsertSchoolClass(w http.ResponseWriter, r *http.Request) {
	var entity model.SchoolClass
	_ = json.NewDecoder(r.Body).Decode(&entity)

	id, err := repository.InsertSchoolClass(entity)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		return
	}

	jsonCreatedResponse(w, id)
}

// UpdateSchoolClass ...
func UpdateSchoolClass(w http.ResponseWriter, r *http.Request) {
	var entity model.SchoolClass
	_ = json.NewDecoder(r.Body).Decode(&entity)

	if err := repository.UpdateSchoolClass(entity); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		return
	}
}

// DeleteSchoolClass ...
func DeleteSchoolClass(w http.ResponseWriter, r *http.Request) {
	// TODO
}
