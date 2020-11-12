package repository

import (
	"log"

	"github.com/felipesulzbach/exemplo-api-rest/app/src/model"
	"github.com/felipesulzbach/exemplo-api-rest/app/src/util"

)

// FindAllCourse ...
func FindAllCourse() ([]*model.Course, error) {
	objectMap, err := getAll("SELECT * FROM fs_auto.course")
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
	db, err := newDB()
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	row := db.QueryRow("SELECT * FROM fs_auto.course WHERE id=$1", id)

	result := new(model.Course)
	if err := row.Scan(&result.ID, &result.Name, &result.Description, &result.RegistrationDate); err != nil {
		return nil, err
	}

	db.closeDB()
	return result, nil
}

// InsertCourse ...
func InsertCourse(entity model.Course) (*model.Course, error) {
	db, err := newDB()
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	sqlStatement := "INSERT INTO fs_auto.course (name, description, registration_date) VALUES ($1, $2, $3) RETURNING *"
	result := new(model.Course)
	if err := db.QueryRow(sqlStatement, entity.Name, entity.Description, entity.RegistrationDate).Scan(&result.ID, &result.Name, &result.Description, &result.RegistrationDate); err != nil {
		return nil, err
	}

	db.closeDB()
	return result, nil
}

// UpdateCourse ...
func UpdateCourse(entity model.Course) (*model.Course, error) {
	db, err := newDB()
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	sqlStatement := "UPDATE fs_auto.course SET name=$2, description=$3 WHERE id=$1"
	if _, err := db.Exec(sqlStatement, entity.ID, entity.Name, entity.Description); err != nil {
		return nil, err
	}

	row := db.QueryRow("SELECT * FROM fs_auto.course WHERE id=$1", entity.ID)

	result := new(model.Course)
	if err := row.Scan(&result.ID, &result.Name, &result.Description, &result.RegistrationDate); err != nil {
		return nil, err
	}

	db.closeDB()
	return result, nil
}

// DeleteCourse ...
func DeleteCourse(id int64) error {
	db, err := newDB()
	if err != nil {
		log.Panic(err)
		return err
	}

	sqlStatement := "DELETE fs_auto.course WHERE id=$1"
	if _, err := db.Exec(sqlStatement, id); err != nil {
		return err
	}

	db.closeDB()
	return nil
}
