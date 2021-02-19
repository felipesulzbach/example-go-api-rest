package contractteacher

/*
import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/felipesulzbach/example-go-api-rest/src/model"
	"github.com/gorilla/mux"

)

// TeacherContract ...
type TeacherContract struct {
	Entity model.Teacher
}

// ValidateBodyCreate ...
func (contract TeacherContract) ValidateBodyCreate(r *http.Request) (model.Teacher, error) {
	var course model.Teacher
	if err := json.NewDecoder(r.Body).Decode(&contract.Entity); err != nil {
		return course, err
	}

	if err := contract._validateCreate(); err != nil {
		return course, err
	}

	request := model.Teacher{}
	request.Person = contract.Entity.Person
	request.Course = contract.Entity.Course
	request.Person.RegistrationDate = time.Now()

	return request, nil
}

// ValidateBodyUpdate ...
func (contract TeacherContract) ValidateBodyUpdate(r *http.Request) (model.Teacher, error) {
	var course model.Teacher
	if err := json.NewDecoder(r.Body).Decode(&contract.Entity); err != nil {
		return course, err
	}

	if err := contract._validateUpdate(); err != nil {
		return course, err
	}

	request := model.Teacher{}
	request.Person = contract.Entity.Person
	request.Course = contract.Entity.Course

	return request, nil
}

// ValidatePath ...
func (contract TeacherContract) ValidatePath(r *http.Request) (int64, error) {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	return id, nil
}

func (contract TeacherContract) _validateCreate() error {
	var personContract PersonContract
	personContract.Entity = contract.Entity.Person
	if _, err := personContract.ValidateBodyCreate(); err != nil {
		return err
	}

	if err := validateID(contract.Entity.Course.ID); err != nil {
		return err
	}

	return nil
}

func (contract TeacherContract) _validateUpdate() error {
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
*/
