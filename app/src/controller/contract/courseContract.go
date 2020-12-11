package contract

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/felipesulzbach/exemplo-api-rest/app/src/model"
	"github.com/gorilla/mux"

)

// CourseContract ...
type CourseContract struct {
	Entity model.Course
}

// ValidateBodyCreate ...
func (contract CourseContract) ValidateBodyCreate(r *http.Request) (model.Course, error) {
	var entity model.Course
	if err := json.NewDecoder(r.Body).Decode(&contract.Entity); err != nil {
		return entity, err
	}

	if err := contract._validateCreate(); err != nil {
		return entity, err
	}

	request := model.Course{}
	request.Name = contract.Entity.Name
	request.Description = contract.Entity.Description
	request.RegistrationDate = time.Now()

	return request, nil
}

// ValidateBodyUpdate ...
func (contract CourseContract) ValidateBodyUpdate(r *http.Request) (model.Course, error) {
	var entity model.Course
	if err := json.NewDecoder(r.Body).Decode(&contract.Entity); err != nil {
		return entity, err
	}

	if err := contract._validateUpdate(); err != nil {
		return entity, err
	}

	request := model.Course{}
	request.ID = contract.Entity.ID
	request.Name = contract.Entity.Name
	request.Description = contract.Entity.Description

	return request, nil
}

// ValidatePath ...
func (contract CourseContract) ValidatePath(r *http.Request) (int64, error) {
	params := mux.Vars(r)
	//id := r.FormValue("id")
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	return id, nil
}

func (contract CourseContract) _validateCreate() error {
	if err := validateStringSpecialCharacter("Name", contract.Entity.Name); err != nil {
		return err
	}

	if err := validateStringSpecialCharacter("Description", contract.Entity.Description); err != nil {
		return err
	}

	return nil
}

func (contract CourseContract) _validateUpdate() error {
	if err := validateID(contract.Entity.ID); err != nil {
		return err
	}

	if err := contract._validateCreate(); err != nil {
		return err
	}

	return nil
}
