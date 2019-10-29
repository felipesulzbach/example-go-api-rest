package service

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/_dev/exemplo-api-rest/model"
	"github.com/_dev/exemplo-api-rest/model/entity"
	"github.com/_dev/exemplo-api-rest/util"

	"github.com/gorilla/mux"
)

// FindAllClass - Returns total list of registered classes.
func FindAllClass(w http.ResponseWriter, r *http.Request) {
	db, err := model.NewDB()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		log.Panic(err)
		return
	}

	list, err := db.FindAllClass()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		db.CloseDB()
		return
	}

	for _, item := range list {
		log.Println(item.ToString())
	}

	db.CloseDB()
	json.NewEncoder(w).Encode(list)
}

// FindByIDClass - Returns a specific class by ID.
func FindByIDClass(w http.ResponseWriter, r *http.Request) {
	db, err := model.NewDB()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		log.Panic(err)
		return
	}

	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		db.CloseDB()
		return
	}

	entityy, err := db.FindByIDClass(id)
	switch {
	case err == sql.ErrNoRows:
		var errorDesc bytes.Buffer
		errorDesc.WriteString("ERROR: No records found for id=")
		errorDesc.WriteString(strconv.FormatInt(id, 10))
		log.Println(errorDesc.String())
		json.NewEncoder(w).Encode(errorDesc.String())
		db.CloseDB()
		return
	case err != nil:
		log.Panic(err)
		db.CloseDB()
		return
	default:
	}

	log.Println(entityy.ToString())
	db.CloseDB()
	json.NewEncoder(w).Encode(entityy)
}

// InsertClass - Inserts a new class record in the data base.
func InsertClass(w http.ResponseWriter, r *http.Request) {
	db, err := model.NewDB()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		log.Panic(err)
		return
	}

	id, err := db.NextIDClass()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		db.CloseDB()
		return
	}

	var entityy entity.Class
	_ = json.NewDecoder(r.Body).Decode(&entityy)
	entityy.ID = id

	idReturned, err := db.InsertClass(entityy)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		db.CloseDB()
		return
	}
	db.CloseDB()
	json.NewEncoder(w).Encode(idReturned)
}

// UpdateClass - Updates a base class record.
func UpdateClass(w http.ResponseWriter, r *http.Request) {
	db, err := model.NewDB()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		log.Panic(err)
		return
	}

	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		db.CloseDB()
		return
	}
	courseID, err := strconv.ParseInt(params["courseID"], 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		db.CloseDB()
		return
	}
	startDate := util.StringToTime(params["startDate"])
	endDate := util.StringToTime(params["endDate"])
	registrationDate := util.StringToTime(params["registrationDate"])

	var entityy entity.Class
	entityy.New(id, courseID, startDate, endDate, registrationDate)

	if err = db.UpdateClass(entityy); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		db.CloseDB()
		return
	}
	db.CloseDB()
}

// DeleteClass - Removes a record from the base.
func DeleteClass(w http.ResponseWriter, r *http.Request) {
	db, err := model.NewDB()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		log.Panic(err)
		return
	}

	if err := Delete(w, r, db, "class", "id", "id"); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		db.CloseDB()
		return
	}
	db.CloseDB()
}
