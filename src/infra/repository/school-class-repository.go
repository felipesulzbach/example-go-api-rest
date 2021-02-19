package repository

import (
	"log"

	"github.com/felipesulzbach/example-go-api-rest/src/infra/model"
	"github.com/felipesulzbach/example-go-api-rest/src/domain/util"

)

// FindAllSchoolClass ...
func FindAllSchoolClass() ([]*model.SchoolClass, error) {
	objectMap, err := getAll("school_class")
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	result := make([]*model.SchoolClass, 0)
	for _, object := range objectMap {
		item := new(model.SchoolClass)

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

// FindByIDSchoolClass ...
func FindByIDSchoolClass(id int64) (*model.SchoolClass, error) {
	object, err := getByID("school_class", id)
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	objectJSON, err := util.Serializer(object)
	if err != nil {
		return nil, err
	}

	item := new(model.SchoolClass)
	err = util.Unserializer(objectJSON, &item)
	if err != nil {
		return nil, err
	}

	return item, nil
}

// InsertSchoolClass ...
func InsertSchoolClass(entity model.SchoolClass) (*model.SchoolClass, error) {
	item := new(model.SchoolClass)
	id, err := create(item, entity)
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	entity.ID = id
	object, err := FindByIDSchoolClass(id)
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

// UpdateSchoolClass ...
func UpdateSchoolClass(entity model.SchoolClass) (*model.SchoolClass, error) {
	entityUpdate, err := FindByIDSchoolClass(entity.ID)
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	entityUpdate.Course = entity.Course
	entityUpdate.StartDate = entity.StartDate
	entityUpdate.EndDate = entity.EndDate

	if err := update(entityUpdate); err != nil {
		log.Panic(err)
		return nil, err
	}

	result, err := FindByIDSchoolClass(entity.ID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteSchoolClass ...
func DeleteSchoolClass(id int64) error {
	err := delete("school_class", id)
	if err != nil {
		log.Panic(err)
		return err
	}

	return nil
}
