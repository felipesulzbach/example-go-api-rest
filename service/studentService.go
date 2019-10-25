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

	list, err := model.FindAllStudent(db)
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
		model.CloseDB(db)
		return
	}

	entityy, err := model.FindByIDStudent(db, id)
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

// InsertStudent - Inserts a new student record in the data base.
func InsertStudent(w http.ResponseWriter, r *http.Request) {
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
	registrationDate := util.StringToTime(params["registrationDate"])
	classID, err := strconv.ParseInt(params["classID"], 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		model.CloseDB(db)
		return
	}

	entityyClass, err := model.FindByIDClass(db, classID)
	switch {
	case err == sql.ErrNoRows:
		var errorDesc bytes.Buffer
		errorDesc.WriteString("ERROR: No records found for classID=")
		errorDesc.WriteString(strconv.FormatInt(classID, 10))
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
	person.New(personID, name, cpf, cellPhone, city, zipCode, address, registrationDate)
	personIDReturn, err := model.InsertPerson(db, person)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		model.CloseDB(db)
		return
	}

	var entityy entity.Student
	entityy.New(personIDReturn, entityyClass.ID)
	idReturned, err := model.InsertStudent(db, entityy)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		model.CloseDB(db)
		return
	}
	model.CloseDB(db)
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
	personID, err := strconv.ParseInt(params["personID"], 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		model.CloseDB(db)
		return
	}
	classID, err := strconv.ParseInt(params["classID"], 10, 64)
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

	var entityy entity.Student
	entityy.New(personID, classID)
	var entityyperson entity.Person
	entityyperson.New(personID, name, cpf, cellPhone, city, zipCode, address, registrationDate)

	if err = model.UpdateStudent(db, entityy, entityyperson); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		model.CloseDB(db)
		return
	}
	model.CloseDB(db)
}

// DeleteStudent - Removes a record from the base.
func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	db, err := model.NewDB()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		log.Panic(err)
		return
	}

	if err := Delete(w, r, db, "student", "id_person", "personID"); err != nil {
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
