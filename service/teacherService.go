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

// FindAllTeacher - Returns total list of registered teachers.
func FindAllTeacher(w http.ResponseWriter, r *http.Request) {
	db, err := model.NewDB()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		log.Panic(err)
		return
	}

	list, err := model.FindAllTeacher(db)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		model.CloseDB(db)
		return
	}

	for _, item := range list {
		log.Println(item.ToString())
	}

	model.CloseDB(db)
	json.NewEncoder(w).Encode(list)
}

// FindByIDTeacher - Returns a specific teacher by ID.
func FindByIDTeacher(w http.ResponseWriter, r *http.Request) {
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
		model.CloseDB(db)
		return
	}

	entityy, err := model.FindByIDTeacher(db, id)
	switch {
	case err == sql.ErrNoRows:
		var errorDesc bytes.Buffer
		errorDesc.WriteString("ERROR: No records found for id=")
		errorDesc.WriteString(strconv.FormatInt(id, 10))
		log.Println(errorDesc.String())
		json.NewEncoder(w).Encode(errorDesc.String())
		model.CloseDB(db)
		return
	case err != nil:
		log.Panic(err)
		model.CloseDB(db)
		return
	default:
	}

	log.Println(entityy.ToString())
	model.CloseDB(db)
	json.NewEncoder(w).Encode(entityy)
}

// InsertTeacher - Inserts a new class record in the data base.
func InsertTeacher(w http.ResponseWriter, r *http.Request) {
	db, err := model.NewDB()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		log.Panic(err)
		return
	}

	personID, err := model.NextIDPerson(db)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		model.CloseDB(db)
		return
	}

	params := mux.Vars(r)
	name := params["name"]
	cpf := params["cpf"]
	cellPhone := params["cellPhone"]
	city := params["city"]
	zipCode := params["zipCode"]
	address := params["address"]
	dataCadastro := util.StringToTime(params["registrationDate"])
	courseID, err := strconv.ParseInt(params["courseID"], 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		model.CloseDB(db)
		return
	}

	entityyCourse, err := model.FindByIDCourse(db, courseID)
	switch {
	case err == sql.ErrNoRows:
		var errorDesc bytes.Buffer
		errorDesc.WriteString("ERROR: No records found for courseID=")
		errorDesc.WriteString(strconv.FormatInt(courseID, 10))
		log.Println(errorDesc.String())
		json.NewEncoder(w).Encode(errorDesc.String())
		model.CloseDB(db)
		return
	case err != nil:
		log.Panic(err)
		model.CloseDB(db)
		return
	default:
	}

	var person entity.Person
	person.New(personID, name, cpf, cellPhone, city, zipCode, address, dataCadastro)
	personIDRetorno, err := model.InsertPerson(db, person)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		model.CloseDB(db)
		return
	}

	var entityy entity.Teacher
	entityy.New(personIDRetorno, entityyCourse.ID)
	idReturned, err := model.InsertTeacher(db, entityy)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		model.CloseDB(db)
		return
	}
	model.CloseDB(db)
	json.NewEncoder(w).Encode(idReturned)
}

// UpdateTeacher - Updates a base teacher record.
func UpdateTeacher(w http.ResponseWriter, r *http.Request) {
	db, err := model.NewDB()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		log.Panic(err)
		return
	}

	params := mux.Vars(r)
	personID, err := strconv.ParseInt(params["personID"], 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		model.CloseDB(db)
		return
	}
	courseID, err := strconv.ParseInt(params["courseID"], 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		model.CloseDB(db)
		return
	}
	name := params["name"]
	cpf := params["cpf"]
	cellPhone := params["cellPhone"]
	city := params["city"]
	zipCode := params["zipCode"]
	address := params["address"]
	registrationDate := util.StringToTime(params["registrationDate"])

	var entityy entity.Teacher
	entityy.New(personID, courseID)
	var entityyperson entity.Person
	entityyperson.New(personID, name, cpf, cellPhone, city, zipCode, address, registrationDate)

	if err = model.UpdateTeacher(db, entityy, entityyperson); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		model.CloseDB(db)
		return
	}
	model.CloseDB(db)
}

// DeleteTeacher - Removes a record from the base.
func DeleteTeacher(w http.ResponseWriter, r *http.Request) {
	db, err := model.NewDB()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		log.Panic(err)
		return
	}

	if err := Delete(w, r, db, "teacher", "id_person", "personID"); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		model.CloseDB(db)
		return
	}

	if err := Delete(w, r, db, "person", "id", "personID"); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		model.CloseDB(db)
		return
	}
	model.CloseDB(db)
}
