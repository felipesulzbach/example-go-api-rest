package model

import (
	"github.com/_dev/exemplo-api-rest/model/entity"
)

// NextIDCourse - Returns the next ID.
func (db *DB) NextIDCourse() (int64, error) {
	row := db.QueryRow("SELECT (MAX(id) + 1) FROM GO_TST.course")

	var id int64
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// FindAllCourse - Returns total list of registered courses.
func (db *DB) FindAllCourse() ([]*entity.Course, error) {
	rows, err := db.Query("SELECT * FROM GO_TST.course")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make([]*entity.Course, 0)
	for rows.Next() {
		item := new(entity.Course)
		err := rows.Scan(&item.ID, &item.Name, &item.Description, &item.RegistrationDate)
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

// FindByIDCourse - Returns a specific course by ID.
func (db *DB) FindByIDCourse(id int64) (*entity.Course, error) {
	row := db.QueryRow("SELECT * FROM GO_TST.course WHERE id=$1", id)

	item := new(entity.Course)
	err := row.Scan(&item.ID, &item.Name, &item.Description, &item.RegistrationDate)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// InsertCourse - Inserts a new class record in the data base.
func (db *DB) InsertCourse(entityy entity.Course) (int64, error) {
	sqlStatement := "INSERT INTO GO_TST.course (id, name, description, registration_date) VALUES ($1, $2, $3, $4) RETURNING id"
	var returnedID int64
	err := db.QueryRow(sqlStatement, entityy.ID, entityy.Name, entityy.Description, entityy.RegistrationDate).Scan(&returnedID)
	if err != nil {
		return 0, err
	}

	return returnedID, nil
}

// UpdateCourse - Updates a base class record.
func (db *DB) UpdateCourse(entityy entity.Course) error {
	sqlStatement := "UPDATE GO_TST.course SET name=$2, description=$3, registration_date=$4 WHERE id=$1"
	_, err := db.Exec(sqlStatement, entityy.ID, entityy.Name, entityy.Description, entityy.RegistrationDate)
	if err != nil {
		return err
	}
	return nil
}
