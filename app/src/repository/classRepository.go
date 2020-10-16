package repository

import (
	"github.com/felipesulzbach/exemplo-api-rest/app/src/model"

)

// NextIDClass - Returns the next ID.
func (db *DB) NextIDClass() (int64, error) {
	row := db.QueryRow("SELECT (MAX(id) + 1) FROM GO_TST.class")

	var id int64
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// FindAllClass - Returns total list of registered classes.
func (db *DB) FindAllClass() ([]*model.Class, error) {
	rows, err := db.Query("SELECT * FROM GO_TST.class")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make([]*model.Class, 0)
	for rows.Next() {
		item := new(model.Class)
		err := rows.Scan(&item.ID, &item.Course.ID, &item.StartDate, &item.EndDate, &item.RegistrationDate)
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
func (db *DB) FindByIDClass(id int64) (*model.Class, error) {
	row := db.QueryRow("SELECT * FROM GO_TST.class WHERE id=$1", id)

	item := new(model.Class)
	err := row.Scan(&item.ID, &item.Course.ID, &item.StartDate, &item.EndDate, &item.RegistrationDate)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// InsertClass - Inserts a new class record in the data base.
func (db *DB) InsertClass(modell model.Class) (int64, error) {
	sqlStatement := "INSERT INTO GO_TST.class (id, course_id, start_date, end_date, registration_date) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	var returnedID int64
	err := db.QueryRow(sqlStatement, modell.ID, modell.Course.ID, modell.StartDate, modell.EndDate, modell.RegistrationDate).Scan(&returnedID)
	if err != nil {
		return 0, err
	}

	return returnedID, nil
}

// UpdateClass - Updates a base class record.
func (db *DB) UpdateClass(modely model.Class) error {
	sqlStatement := "UPDATE GO_TST.class SET course_id=$2, start_date=$3, end_date=$4, registration_date=$5 WHERE id=$1"
	_, err := db.Exec(sqlStatement, modely.ID, modely.Course.ID, modely.StartDate, modely.EndDate, modely.RegistrationDate)
	if err != nil {
		return err
	}
	return nil
}
