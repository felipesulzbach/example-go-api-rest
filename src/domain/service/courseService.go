package service

import (
	"log"

	contract "github.com/felipesulzbach/example-go-api-rest/src/infra/contract/contractcourse"
	"github.com/felipesulzbach/example-go-api-rest/src/infra/model"
	"github.com/felipesulzbach/example-go-api-rest/src/infra/repository"

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
func FindByIDCourse(path contract.Get) (*model.Course, error) {
	result, err := repository.FindByIDCourse(path.ID)
	if err != nil {
		return nil, err
	}

	log.Println(result.ToString())
	return result, nil
}

// InsertCourse ...
func InsertCourse(body contract.Post) (*model.Course, error) {
	entity := *new(model.Course)
	entity.Name = body.Name
	entity.Description = body.Description

	result, err := repository.InsertCourse(entity)
	if err != nil {
		return nil, err
	}

	log.Println(result.ToString())
	return result, nil
}

// UpdateCourse ...
func UpdateCourse(body contract.Put) (*model.Course, error) {
	entity := *new(model.Course)
	entity.ID = body.ID
	entity.Name = body.Name
	entity.Description = body.Description

	result, err := repository.UpdateCourse(entity)
	if err != nil {
		return nil, err
	}

	log.Println(result.ToString())
	return result, nil
}

// DeleteCourse ...
func DeleteCourse(path contract.Get) error {
	if err := repository.DeleteCourse(path.ID); err != nil {
		return err
	}

	return nil
}
