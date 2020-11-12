package service

import (
	"log"
	"time"

	"github.com/felipesulzbach/exemplo-api-rest/app/src/model"
	"github.com/felipesulzbach/exemplo-api-rest/app/src/repository"

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
	entity, err := repository.FindByIDSchoolClass(id)
	if err != nil {
		return nil, err
	}

	log.Println(entity.ToString())
	return entity, nil
}

// InsertSchoolClass ...
func InsertSchoolClass(entity model.SchoolClass) (int64, error) {
	entity.RegistrationDate = time.Now()
	id, err := repository.InsertSchoolClass(entity)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// UpdateSchoolClass ...
func UpdateSchoolClass(entity model.SchoolClass) error {
	if err := repository.UpdateSchoolClass(entity); err != nil {
		return err
	}

	return nil
}

// DeleteSchoolClass ...
func DeleteSchoolClass() {
	// TODO
}
