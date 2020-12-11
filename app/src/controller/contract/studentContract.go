package contract

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/felipesulzbach/exemplo-api-rest/app/src/model"
	"github.com/gorilla/mux"

)

// StudentContract ...
type StudentContract struct {
	Entity model.Student
}

// ValidateBodyCreate ...
func (contract StudentContract) ValidateBodyCreate(r *http.Request) (model.Student, error) {
	var entity model.Student
	if err := json.NewDecoder(r.Body).Decode(&contract.Entity); err != nil {
		return entity, err
	}

	if err := contract._validateCreate(); err != nil {
		return entity, err
	}

	request := model.Student{}
	request.Person = contract.Entity.Person
	request.SchoolClass = contract.Entity.SchoolClass
	request.Person.RegistrationDate = time.Now()

	return request, nil
}

// ValidateBodyUpdate ...
func (contract StudentContract) ValidateBodyUpdate(r *http.Request) (model.Student, error) {
	var entity model.Student
	if err := json.NewDecoder(r.Body).Decode(&contract.Entity); err != nil {
		return entity, err
	}

	if err := contract._validateUpdate(); err != nil {
		return entity, err
	}

	request := model.Student{}
	request.Person = contract.Entity.Person
	request.SchoolClass = contract.Entity.SchoolClass

	return request, nil
}

// ValidatePath ...
func (contract StudentContract) ValidatePath(r *http.Request) (int64, error) {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	return id, nil
}

func (contract StudentContract) _validateCreate() error {
	var personContract PersonContract
	personContract.Entity = contract.Entity.Person
	if _, err := personContract.ValidateBodyCreate(); err != nil {
		return err
	}

	if err := validateID(contract.Entity.SchoolClass.ID); err != nil {
		return err
	}

	return nil
}

func (contract StudentContract) _validateUpdate() error {
	var personContract PersonContract
	personContract.Entity = contract.Entity.Person
	if _, err := personContract.ValidateBodyUpdate(); err != nil {
		return err
	}

	if err := contract._validateCreate(); err != nil {
		return err
	}

	return nil
}
