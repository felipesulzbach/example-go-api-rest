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
func InsertCourse(entity model.Course) (int64, error) {
	id, err := repository.InsertCourse(entity)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// UpdateCourse ...
func UpdateCourse(entity model.Course) error {
	if err := repository.UpdateCourse(entity); err != nil {
		return err
	}
	return nil
}

// DeleteCourse ...
func DeleteCourse() {
	// TODO
}
