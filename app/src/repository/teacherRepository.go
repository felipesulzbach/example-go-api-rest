package repository

import (
	"github.com/felipesulzbach/exemplo-api-rest/app/src/model"

)

// FindAllTeacher - Returns total list of registered teachers.
func (db *DB) FindAllTeacher() ([]*model.Teacher, error) {
	rows, err := db.Query("SELECT * FROM GO_TST.teacher")
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
	return list, nil
}

// FindByIDTeacher - Returns a specific teacher by ID.
func (db *DB) FindByIDTeacher(id int64) (*model.Teacher, error) {
	row := db.QueryRow("SELECT * FROM GO_TST.teacher WHERE id=$1", id)

	item := new(model.Teacher)
	err := row.Scan(&item.Person.ID, &item.Course.ID)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// InsertTeacher - Inserts a new class record in the data base.
func (db *DB) InsertTeacher(modely model.Teacher) (int64, error) {
	sqlStatement := "INSERT INTO GO_TST.teacher (person_id, course_id) VALUES ($1, $2) RETURNING person_id"
	var returnedID int64
	err := db.QueryRow(sqlStatement, modely.Person.ID, modely.Course.ID).Scan(&returnedID)
	if err != nil {
		return 0, err
	}

	return returnedID, nil
}

// UpdateTeacher - Updates a base teacher record.
func (db *DB) UpdateTeacher(modely model.Teacher) error {
	sqlStatement := "UPDATE GO_TST.teacher SET course_id=$2 WHERE person_id=$1"
	_, err := db.Exec(sqlStatement, modely.Person.ID, modely.Course.ID)
	if err != nil {
		return err
	}

	if err = db.UpdatePerson(modely.Person); err != nil {
		return err
	}
	return nil
}
