package service

import (
	"log"
	"time"

	"github.com/felipesulzbach/exemplo-api-rest/app/src/model"
	"github.com/felipesulzbach/exemplo-api-rest/app/src/repository"

)

// FindAllStudent ...
func FindAllStudent() ([]*model.Student, error) {
	result, err := repository.FindAllStudent()
	if err != nil {
		return nil, err
	}

	for _, item := range result {
		log.Println(item.ToString())
	}

	return result, nil
}

// FindByIDStudent ...
func FindByIDStudent(id int64) (*model.Student, error) {
	result, err := repository.FindByIDStudent(id)
	if err != nil {
		return nil, err
	}

	log.Println(result.ToString())
	return result, nil
}

// InsertStudent ...
func InsertStudent(entity model.Student) (int64, error) {
	entity.Person.RegistrationDate = time.Now()
	id, err := repository.InsertStudent(entity)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// UpdateStudent ...
func UpdateStudent(entity model.Student) error {
	if err := repository.UpdateStudent(entity); err != nil {
		return err
	}

	return nil
}

// DeleteStudent ...
func DeleteStudent() {
	// TODO
}
