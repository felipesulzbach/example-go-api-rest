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

func getAllCourse(w http.ResponseWriter, r *http.Request) {
	//	list, err := env.db.FindAllCourse()
	response, err := service.FindAllCourse()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		return
	}

	jsonOkResponse(w, response)
}

func getByIDCourse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	//id := r.FormValue("id")
	id, _ := strconv.ParseInt(params["id"], 10, 64)

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
	var entity model.Course
	_ = json.NewDecoder(r.Body).Decode(&entity)

	id, err := service.InsertCourse(entity)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		return
	}

	jsonCreatedResponse(w, id)
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	var entity model.Course
	_ = json.NewDecoder(r.Body).Decode(&entity)

	if err := service.UpdateCourse(entity); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		return
	}
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	// TODO
}
