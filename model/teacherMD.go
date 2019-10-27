package model

import (
	"github.com/_dev/exemplo-api-rest/model/entity"
)

// FindAllTeacher - Returns total list of registered teachers.
func (db *DB) FindAllTeacher() ([]*entity.Teacher, error) {
	rows, err := db.Query("SELECT * FROM GO_TST.teacher")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make([]*entity.Teacher, 0)
	for rows.Next() {
		item := new(entity.Teacher)
		err := rows.Scan(&item.PersonID, &item.CourseID)
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
func (db *DB) FindByIDTeacher(id int64) (*entity.Teacher, error) {
	row := db.QueryRow("SELECT * FROM GO_TST.teacher WHERE id=$1", id)

	item := new(entity.Teacher)
	err := row.Scan(&item.PersonID, &item.CourseID)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// InsertTeacher - Inserts a new class record in the data base.
func (db *DB) InsertTeacher(entityy entity.Teacher) (int64, error) {
	sqlStatement := "INSERT INTO GO_TST.teacher (person_id, course_id) VALUES ($1, $2) RETURNING person_id"

	var returnedID int64
	err := db.QueryRow(sqlStatement, entityy.PersonID, entityy.CourseID).Scan(&returnedID)
	if err != nil {
		return 0, err
	}

	return returnedID, nil
}

// UpdateTeacher - Updates a base teacher record.
func (db *DB) UpdateTeacher(entityy entity.Teacher, person entity.Person) error {
	sqlStatement := "UPDATE GO_TST.teacher SET course_id=$2 WHERE person_id=$1"
	_, err := db.Exec(sqlStatement, entityy.PersonID, entityy.CourseID)
	if err != nil {
		return err
	}
	if err = db.UpdatePerson(person); err != nil {
		return err
	}
	return nil
}
