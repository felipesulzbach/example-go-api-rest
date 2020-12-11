package contract

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/felipesulzbach/exemplo-api-rest/app/src/model"
	"github.com/gorilla/mux"

)

// SchoolClassContract ...
type SchoolClassContract struct {
	Entity model.SchoolClass
}

// ValidateBodyCreate ...
func (contract SchoolClassContract) ValidateBodyCreate(r *http.Request) (model.SchoolClass, error) {
	var entity model.SchoolClass
	if err := json.NewDecoder(r.Body).Decode(&contract.Entity); err != nil {
		return entity, err
	}

	if err := contract._validateCreate(); err != nil {
		return entity, err
	}

	request := model.SchoolClass{}
	request.Course = contract.Entity.Course
	request.StartDate = contract.Entity.StartDate
	request.EndDate = contract.Entity.EndDate
	request.RegistrationDate = time.Now()

	return request, nil
}

// ValidateBodyUpdate ...
func (contract SchoolClassContract) ValidateBodyUpdate(r *http.Request) (model.SchoolClass, error) {
	var entity model.SchoolClass
	if err := json.NewDecoder(r.Body).Decode(&contract.Entity); err != nil {
		return entity, err
	}

	if err := contract._validateUpdate(); err != nil {
		return entity, err
	}

	request := model.SchoolClass{}
	request.Course = contract.Entity.Course
	request.StartDate = contract.Entity.StartDate
	request.EndDate = contract.Entity.EndDate

	return request, nil
}

// ValidatePath ...
func (contract SchoolClassContract) ValidatePath(r *http.Request) (int64, error) {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	return id, nil
}

func (contract SchoolClassContract) _validateCreate() error {
	if err := validateID(contract.Entity.Course.ID); err != nil {
		return err
	}

	if err := validateDateIsDateTimeEmpty("StartDate", contract.Entity.StartDate); err != nil {
		return err
	}

	if err := validateDateIsDateTimeEmpty("EndDate", contract.Entity.EndDate); err != nil {
		return err
	}

	if err := validadeStarDateAfterEndDate(contract.Entity.StartDate, contract.Entity.EndDate); err != nil {
		return err
	}

	return nil
}

func (contract SchoolClassContract) _validateUpdate() error {
	if err := validateID(contract.Entity.ID); err != nil {
		return err
	}

	if err := contract._validateCreate(); err != nil {
		return err
	}

	return nil
}
