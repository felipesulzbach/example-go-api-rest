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

// FindAllStudent - Returns total list of registered students.
func FindAllStudent(w http.ResponseWriter, r *http.Request) {
	db, err := model.NewDB()
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

	entityy, err := db.FindByIDStudent(id)
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

// InsertStudent - Inserts a new student record in the data base.
func InsertStudent(w http.ResponseWriter, r *http.Request) {
	db, err := model.NewDB()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		log.Panic(err)
		return
	}

	var entityy entity.Student
	_ = json.NewDecoder(r.Body).Decode(&entityy)

	id, err := db.InsertPerson(entityy.Person)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		db.CloseDB()
		return
	}
	entityy.Person.ID = id;

	idClass, err := db.InsertClass(entityy.Class)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		db.CloseDB()
		return
	}
	entityy.Class.ID = idClass;

	idReturned, err := db.InsertStudent(entityy)
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
	classID, err := strconv.ParseInt(params["classID"], 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		db.CloseDB()
		return
	}
	name := params["name"]
	cpf := params["cpf"]
	cellPhone := params["cellPhone"]
	city := params["city"]
	zipCode := params["zipCode"]
	address := params["address"]
	registrationDate := util.StringToTime(params["registrationDate"])

	var entityy entity.Student
	entityy.New(id, classID)
	var entityyperson entity.Person
	entityyperson.New(id, name, cpf, cellPhone, city, zipCode, address, registrationDate)

	if err = db.UpdateStudent(entityy, entityyperson); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		db.CloseDB()
		return
	}
	db.CloseDB()
}

// DeleteStudent - Removes a record from the base.
func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	db, err := model.NewDB()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		log.Panic(err)
		return
	}

	if err := Delete(w, r, db, "student", "id_person", "id"); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		db.CloseDB()
		return
	}

	if err := Delete(w, r, db, "person", "id", "id"); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		db.CloseDB()
		return
	}
	db.CloseDB()
}
