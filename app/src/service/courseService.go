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

// FindAllCourse ...
func FindAllCourse(w http.ResponseWriter, r *http.Request) {
	//	list, err := env.db.FindAllCourse()
	list, err := repository.FindAllCourse()
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

// FindByIDCourse ...
func FindByIDCourse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	//id := r.FormValue("id")
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		return
	}

	entity, err := repository.FindByIDCourse(id)
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

// InsertCourse ...
func InsertCourse(w http.ResponseWriter, r *http.Request) {
	var entity model.Course
	_ = json.NewDecoder(r.Body).Decode(&entity)

	idReturned, err := repository.InsertCourse(entity)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		return
	}
	jsonCreatedResponse(w, idReturned)
}

// UpdateCourse ...
func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	var entity model.Course
	_ = json.NewDecoder(r.Body).Decode(&entity)

	if err := repository.UpdateCourse(entity); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		return
	}
}

// DeleteCourse ...
func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	// TODO
}
