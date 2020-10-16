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

// FindAllStudent - Returns total list of registered students.
func FindAllStudent(w http.ResponseWriter, r *http.Request) {
	db, err := repository.NewDB()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		log.Panic(err)
		return
	}

	list, err := db.FindAllStudent()
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

// FindByIDStudent - Returns a specific student by ID.
func FindByIDStudent(w http.ResponseWriter, r *http.Request) {
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

	modely, err := db.FindByIDStudent(id)
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

// InsertStudent - Inserts a new student record in the data base.
func InsertStudent(w http.ResponseWriter, r *http.Request) {
	db, err := repository.NewDB()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		log.Panic(err)
		return
	}

	var modely model.Student
	_ = json.NewDecoder(r.Body).Decode(&modely)

	id, err := db.InsertPerson(modely.Person)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		db.CloseDB()
		return
	}
	modely.Person.ID = id

	modelyClass, err := db.FindByIDCourse(modely.Class.ID)
	switch {
	case err == sql.ErrNoRows:
		idClass, err := db.InsertClass(modely.Class)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			log.Panic(err)
			db.CloseDB()
			return
		}
		modely.Class.ID = idClass
	case err != nil:
		log.Panic(err)
		db.CloseDB()
		return
	default:
		modely.Class.ID = modelyClass.ID
	}

	idReturned, err := db.InsertStudent(modely)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		db.CloseDB()
		return
	}
	db.CloseDB()
	json.NewEncoder(w).Encode(idReturned)
}

// UpdateStudent - Updates a base student record.
func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	db, err := repository.NewDB()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		log.Panic(err)
		return
	}

	var modely model.Student
	_ = json.NewDecoder(r.Body).Decode(&modely)

	if err = db.UpdateStudent(modely); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		db.CloseDB()
		return
	}
	db.CloseDB()
}

// DeleteStudent - Removes a record from the base.
func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	db, err := repository.NewDB()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		log.Panic(err)
		return
	}

	delPO := deletePO{"student", "id_person", "id"}
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
