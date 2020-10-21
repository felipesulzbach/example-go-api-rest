package repository

import (
	"log"

	"github.com/felipesulzbach/exemplo-api-rest/app/src/model"

)

// FindAllTeacher ...
func FindAllTeacher() ([]*model.Teacher, error) {
	db, err := newDB()
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	rows, err := db.Query("SELECT * FROM fs_auto.teacher")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make([]*model.Teacher, 0)
	for rows.Next() {
		item := new(model.Teacher)
		err := rows.Scan(&item.Person.ID, &item.Course.ID)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	db.closeDB()
	return list, nil
}

// FindByIDTeacher ...
func FindByIDTeacher(id int64) (*model.Teacher, error) {
	db, err := newDB()
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	row := db.QueryRow("SELECT * FROM fs_auto.teacher WHERE id=$1", id)
	item := new(model.Teacher)
	if err := row.Scan(&item.Person.ID, &item.Course.ID); err != nil {
		return nil, err
	}

	db.closeDB()
	return item, nil
}

// InsertTeacher ...
func InsertTeacher(entity model.Teacher) (int64, error) {
	db, err := newDB()
	if err != nil {
		log.Panic(err)
		return 0, err
	}

	sqlStatement := "INSERT INTO fs_auto.teacher (course_id) VALUES ($1) RETURNING id"
	var returnedID int64
	if err := db.QueryRow(sqlStatement, entity.Course.ID).Scan(&returnedID); err != nil {
		return 0, err
	}

	db.closeDB()
	return returnedID, nil
}

// UpdateTeacher ...
func UpdateTeacher(entity model.Teacher) error {
	db, err := newDB()
	if err != nil {
		log.Panic(err)
		return err
	}

	sqlStatement := "UPDATE fs_auto.teacher SET course_id=$1 WHERE id=$1"
	if _, err := db.Exec(sqlStatement, entity.Course.ID); err != nil {
		return err
	}

	db.closeDB()
	return nil
}
