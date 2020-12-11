package repository

import (
	"log"

	"github.com/felipesulzbach/exemplo-api-rest/app/src/model"
	"github.com/felipesulzbach/exemplo-api-rest/app/src/util"

)

// FindAllStudent ...
func FindAllStudent() ([]*model.Student, error) {
	objectMap, err := getAll("student")
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	result := make([]*model.Student, 0)
	for _, object := range objectMap {
		item := new(model.Student)

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

// FindByIDStudent ...
func FindByIDStudent(id int64) (*model.Student, error) {
	object, err := getByID("student", id)
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	objectJSON, err := util.Serializer(object)
	if err != nil {
		return nil, err
	}

	item := new(model.Student)
	err = util.Unserializer(objectJSON, &item)
	if err != nil {
		return nil, err
	}

	return item, nil
}

// InsertStudent ...
func InsertStudent(entity model.Student) (*model.Student, error) {
	item := new(model.Student)
	id, err := create(item, entity)
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	entity.Person.ID = id
	object, err := FindByIDStudent(id)
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

// UpdateStudent ...
func UpdateStudent(entity model.Student) (*model.Student, error) {
	entityUpdate, err := FindByIDStudent(entity.Person.ID)
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	entityUpdate.Person = entity.Person
	entityUpdate.SchoolClass = entity.SchoolClass

	if err := update(entityUpdate); err != nil {
		log.Panic(err)
		return nil, err
	}

	result, err := FindByIDStudent(entity.Person.ID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteStudent ...
func DeleteStudent(id int64) error {
	err := delete("student", id)
	if err != nil {
		log.Panic(err)
		return err
	}

	return nil
}
