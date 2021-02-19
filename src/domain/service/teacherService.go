package service

import (
	"log"

	"github.com/felipesulzbach/example-go-api-rest/src/infra/model"
	"github.com/felipesulzbach/example-go-api-rest/src/infra/repository"

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
func InsertTeacher(entity model.Teacher) (*model.Teacher, error) {
	result, err := repository.InsertTeacher(entity)
	if err != nil {
		return nil, err
	}

	log.Println(result.ToString())
	return result, nil
}

// UpdateTeacher ...
func UpdateTeacher(entity model.Teacher) (*model.Teacher, error) {
	result, err := repository.UpdateTeacher(entity)
	if err != nil {
		return nil, err
	}

	log.Println(result.ToString())
	return result, nil
}

// DeleteTeacher ...
func DeleteTeacher(id int64) error {
	if err := repository.DeleteTeacher(id); err != nil {
		return err
	}

	return nil
}
