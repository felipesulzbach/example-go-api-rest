package service

import (
	"log"

	"github.com/felipesulzbach/example-go-api-rest/src/infra/model"
	"github.com/felipesulzbach/example-go-api-rest/src/infra/repository"

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
func InsertStudent(entity model.Student) (*model.Student, error) {
	result, err := repository.InsertStudent(entity)
	if err != nil {
		return nil, err
	}

	log.Println(result.ToString())
	return result, nil
}

// UpdateStudent ...
func UpdateStudent(entity model.Student) (*model.Student, error) {
	result, err := repository.UpdateStudent(entity)
	if err != nil {
		return nil, err
	}

	log.Println(result.ToString())
	return result, nil
}

// DeleteStudent ...
func DeleteStudent(id int64) error {
	if err := repository.DeleteStudent(id); err != nil {
		return err
	}

	return nil
}
