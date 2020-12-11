package repository

import (
	"log"

	"github.com/felipesulzbach/exemplo-api-rest/app/src/model"
	"github.com/felipesulzbach/exemplo-api-rest/app/src/util"

)

// FindAllTeacher ...
func FindAllTeacher() ([]*model.Teacher, error) {
	objectMap, err := getAll("teacher")
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	result := make([]*model.Teacher, 0)
	for _, object := range objectMap {
		item := new(model.Teacher)

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

// FindByIDTeacher ...
func FindByIDTeacher(id int64) (*model.Teacher, error) {
	object, err := getByID("teacher", id)
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	objectJSON, err := util.Serializer(object)
	if err != nil {
		return nil, err
	}

	item := new(model.Teacher)
	err = util.Unserializer(objectJSON, &item)
	if err != nil {
		return nil, err
	}

	return item, nil
}

// InsertTeacher ...
func InsertTeacher(entity model.Teacher) (*model.Teacher, error) {
	item := new(model.Teacher)
	id, err := create(item, entity)
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	entity.Person.ID = id
	object, err := FindByIDTeacher(id)
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

// UpdateTeacher ...
func UpdateTeacher(entity model.Teacher) (*model.Teacher, error) {
	entityUpdate, err := FindByIDTeacher(entity.Person.ID)
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	entityUpdate.Person = entity.Person
	entityUpdate.Course = entity.Course

	if err := update(entityUpdate); err != nil {
		log.Panic(err)
		return nil, err
	}

	result, err := FindByIDTeacher(entity.Person.ID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteTeacher ...
func DeleteTeacher(id int64) error {
	err := delete("teacher", id)
	if err != nil {
		log.Panic(err)
		return err
	}

	return nil
}
