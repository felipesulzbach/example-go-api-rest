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

func getAllSchoolClass(w http.ResponseWriter, r *http.Request) {
	response, err := service.FindAllSchoolClass()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		return
	}

	jsonOkResponse(w, response)
}

func getByIDSchoolClass(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		return
	}

	response, err := service.FindByIDSchoolClass(id)
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

func insertSchoolClass(w http.ResponseWriter, r *http.Request) {
	var entity model.SchoolClass
	_ = json.NewDecoder(r.Body).Decode(&entity)

	id, err := service.InsertSchoolClass(entity)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		return
	}

	jsonCreatedResponse(w, id)
}

func updateSchoolClass(w http.ResponseWriter, r *http.Request) {
	var entity model.SchoolClass
	_ = json.NewDecoder(r.Body).Decode(&entity)

	if err := service.UpdateSchoolClass(entity); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		return
	}
}

func deleteSchoolClass(w http.ResponseWriter, r *http.Request) {
	// TODO
}
