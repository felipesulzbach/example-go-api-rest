package service

import (
	"log"
	"time"

	"github.com/felipesulzbach/exemplo-api-rest/app/src/model"
	"github.com/felipesulzbach/exemplo-api-rest/app/src/repository"

)

// FindAllTeacher ...
func FindAllTeacher() ([]*model.Teacher, error) {
	result, err := repository.FindAllTeacher()
	if err != nil {
		return nil, err
	}

	for _, item := range result {
		log.Println(item.ToString())
	}

	return result, nil
}

// FindByIDTeacher ...
func FindByIDTeacher(id int64) (*model.Teacher, error) {
	result, err := repository.FindByIDTeacher(id)
	if err != nil {
		return nil, err
	}

	log.Println(result.ToString())
	return result, nil
}

// InsertTeacher ...
func InsertTeacher(entity model.Teacher) (int64, error) {
	entity.Person.RegistrationDate = time.Now()
	id, err := repository.InsertTeacher(entity)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// UpdateTeacher ...
func UpdateTeacher(entity model.Teacher) error {
	if err := repository.UpdateTeacher(entity); err != nil {
		return err
	}

	return nil
}

// DeleteTeacher ...
func DeleteTeacher() {
	// TODO
}
