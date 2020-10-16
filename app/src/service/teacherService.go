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

// FindAllTeacher - Returns total list of registered teachers.
func FindAllTeacher(w http.ResponseWriter, r *http.Request) {
	db, err := repository.NewDB()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		log.Panic(err)
		return
	}

	list, err := db.FindAllTeacher()
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

// FindByIDTeacher - Returns a specific teacher by ID.
func FindByIDTeacher(w http.ResponseWriter, r *http.Request) {
	db, err := repository.NewDB()
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

	modely, err := db.FindByIDTeacher(id)
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

	log.Println(modely.ToString())
	db.CloseDB()
	json.NewEncoder(w).Encode(modely)
}

// InsertTeacher - Inserts a new class record in the data base.
func InsertTeacher(w http.ResponseWriter, r *http.Request) {
	db, err := repository.NewDB()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		log.Panic(err)
		return
	}

	var modely model.Teacher
	_ = json.NewDecoder(r.Body).Decode(&modely)

	id, err := db.NextIDPerson()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		db.CloseDB()
		return
	}
	modely.Person.ID = id

	modelyCourse, err := db.FindByIDCourse(modely.Course.ID)
	switch {
	case err == sql.ErrNoRows:
		idCourse, err := db.InsertCourse(modely.Course)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			log.Panic(err)
			db.CloseDB()
			return
		}
		modely.Course.ID = idCourse
	case err != nil:
		log.Panic(err)
		db.CloseDB()
		return
	default:
		modely.Course.ID = modelyCourse.ID
	}

	idReturned, err := db.InsertTeacher(modely)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		db.CloseDB()
		return
	}
	db.CloseDB()
	json.NewEncoder(w).Encode(idReturned)
}

// UpdateTeacher - Updates a base teacher record.
func UpdateTeacher(w http.ResponseWriter, r *http.Request) {
	db, err := repository.NewDB()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		log.Panic(err)
		return
	}

	var modely model.Teacher
	_ = json.NewDecoder(r.Body).Decode(&modely)

	if err = db.UpdateTeacher(modely); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		db.CloseDB()
		return
	}
	db.CloseDB()
}

// DeleteTeacher - Removes a record from the base.
func DeleteTeacher(w http.ResponseWriter, r *http.Request) {
	db, err := repository.NewDB()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		log.Panic(err)
		return
	}

	delPO := deletePO{"teacher", "id_person", "id"}
	if err := delPO.Delete(w, r, db); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		db.CloseDB()
		return
	}

	delPO = deletePO{"person", "id", "id"}
	if err := delPO.Delete(w, r, db); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		db.CloseDB()
		return
	}
	db.CloseDB()
}
