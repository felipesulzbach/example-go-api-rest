package model

import (
	"github.com/_dev/exemplo-api-rest/model/entity"
)

// NextIDClass - Returns the next ID.
func NextIDClass(db *DB) (int64, error) {
	row := db.QueryRow("SELECT (MAX(id) + 1) FROM GO_TST.class")

	var id int64
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// FindAllClass - Returns total list of registered classes.
func FindAllClass(db *DB) ([]*entity.Class, error) {
	rows, err := db.Query("SELECT * FROM GO_TST.class")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make([]*entity.Class, 0)
	for rows.Next() {
		item := new(entity.Class)
		err := rows.Scan(&item.ID, &item.CourseID, &item.StartDate, &item.EndDate, &item.RegistrationDate)
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

// FindByIDClass - Returns a specific class by ID.
func FindByIDClass(db *DB, id int64) (*entity.Class, error) {
	row := db.QueryRow("SELECT * FROM GO_TST.class WHERE id=$1", id)

	item := new(entity.Class)
	err := row.Scan(&item.ID, &item.CourseID, &item.StartDate, &item.EndDate, &item.RegistrationDate)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// InsertClass - Inserts a new class record in the data base.
func InsertClass(db *DB, entityy entity.Class) (int64, error) {
	sqlStatement := "INSERT INTO GO_TST.class (id, course_id, start_date, end_date, registration_date) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	var returnedID int64
	err := db.QueryRow(sqlStatement, entityy.ID, entityy.CourseID, entityy.StartDate, entityy.EndDate, entityy.RegistrationDate).Scan(&returnedID)
	if err != nil {
		return 0, err
	}

	return returnedID, nil
}

// UpdateClass - Updates a base class record.
func UpdateClass(db *DB, entityy entity.Class) error {
	sqlStatement := "UPDATE GO_TST.class SET course_id=$2, start_date=$3, end_date=$4, registration_date=$5 WHERE id=$1"
	_, err := db.Exec(sqlStatement, entityy.ID, entityy.CourseID, entityy.StartDate, entityy.EndDate, entityy.RegistrationDate)
	if err != nil {
		return err
	}
	return nil
}
