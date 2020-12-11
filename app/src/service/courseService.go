package service

import (
	"log"

	"github.com/felipesulzbach/exemplo-api-rest/app/src/model"
	"github.com/felipesulzbach/exemplo-api-rest/app/src/repository"

)

// FindAllCourse ...
func FindAllCourse() ([]*model.Course, error) {
	result, err := repository.FindAllCourse()
	if err != nil {
		return nil, err
	}

	for _, item := range result {
		log.Println(item.ToString())
	}

	return result, nil
}

// FindByIDCourse ...
func FindByIDCourse(id int64) (*model.Course, error) {
	result, err := repository.FindByIDCourse(id)
	if err != nil {
		return nil, err
	}

	log.Println(result.ToString())
	return result, nil
}

// InsertCourse ...
func InsertCourse(entity model.Course) (*model.Course, error) {
	result, err := repository.InsertCourse(entity)
	if err != nil {
		return nil, err
	}

	log.Println(result.ToString())
	return result, nil
}

// UpdateCourse ...
func UpdateCourse(entity model.Course) (*model.Course, error) {
	result, err := repository.UpdateCourse(entity)
	if err != nil {
		return nil, err
	}

	log.Println(result.ToString())
	return result, nil
}

// DeleteCourse ...
func DeleteCourse(id int64) error {
	if err := repository.DeleteCourse(id); err != nil {
		return err
	}

	return nil
}
