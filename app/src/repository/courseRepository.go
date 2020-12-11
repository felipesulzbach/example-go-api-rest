package repository

import (
	"log"
	"time"

	"github.com/felipesulzbach/exemplo-api-rest/app/src/model"
	"github.com/felipesulzbach/exemplo-api-rest/app/src/util"

)

// FindAllCourse ...
func FindAllCourse() ([]*model.Course, error) {
	objectMap, err := getAll("course")
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	result := make([]*model.Course, 0)
	for _, object := range objectMap {
		item := new(model.Course)

		objectJSON, err := util.Serializer(object)
		if err != nil {
			return nil, err
		}

		err = util.Unserializer(objectJSON, &item)
		if err != nil {
			return nil, err
		}

		result = append(result, item)
	}

	return result, nil
}

// FindByIDCourse ...
func FindByIDCourse(id int64) (*model.Course, error) {
	object, err := getByID("course", id)
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	objectJSON, err := util.Serializer(object)
	if err != nil {
		return nil, err
	}

	item := new(model.Course)
	err = util.Unserializer(objectJSON, &item)
	if err != nil {
		return nil, err
	}

	return item, nil
}

// InsertCourse ...
func InsertCourse(entity model.Course) (*model.Course, error) {
	entity.RegistrationDate = time.Now()

	item := new(model.Course)
	id, err := create(item, entity)
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	entity.ID = id
	object, err := FindByIDCourse(id)
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	objectJSON, err := util.Serializer(object)
	if err != nil {
		return nil, err
	}

	err = util.Unserializer(objectJSON, &item)
	if err != nil {
		return nil, err
	}

	return &entity, nil
}

// UpdateCourse ...
func UpdateCourse(entity model.Course) (*model.Course, error) {
	entityUpdate, err := FindByIDCourse(entity.ID)
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	entityUpdate.Name = entity.Name
	entityUpdate.Description = entity.Description

	if err := update(entityUpdate); err != nil {
		log.Panic(err)
		return nil, err
	}

	result, err := FindByIDCourse(entity.ID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteCourse ...
func DeleteCourse(id int64) error {
	err := delete("course", id)
	if err != nil {
		log.Panic(err)
		return err
	}

	return nil
}
