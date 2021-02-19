package service

import (
	"log"

	"github.com/felipesulzbach/example-go-api-rest/src/infra/model"
	"github.com/felipesulzbach/example-go-api-rest/src/infra/repository"

)

// FindAllSchoolClass ...
func FindAllSchoolClass() ([]*model.SchoolClass, error) {
	result, err := repository.FindAllSchoolClass()
	if err != nil {
		return nil, err
	}

	for _, item := range result {
		log.Println(item.ToString())
	}

	return result, nil
}

// FindByIDSchoolClass ...
func FindByIDSchoolClass(id int64) (*model.SchoolClass, error) {
	result, err := repository.FindByIDSchoolClass(id)
	if err != nil {
		return nil, err
	}

	log.Println(result.ToString())
	return result, nil
}

// InsertSchoolClass ...
func InsertSchoolClass(entity model.SchoolClass) (*model.SchoolClass, error) {
	result, err := repository.InsertSchoolClass(entity)
	if err != nil {
		return nil, err
	}

	log.Println(result.ToString())
	return result, nil
}

// UpdateSchoolClass ...
func UpdateSchoolClass(entity model.SchoolClass) (*model.SchoolClass, error) {
	result, err := repository.UpdateSchoolClass(entity)
	if err != nil {
		return nil, err
	}

	log.Println(result.ToString())
	return result, nil
}

// DeleteSchoolClass ...
func DeleteSchoolClass(id int64) error {
	if err := repository.DeleteSchoolClass(id); err != nil {
		return err
	}

	return nil
}
