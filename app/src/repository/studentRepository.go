package repository

import (
	"github.com/felipesulzbach/exemplo-api-rest/app/src/model"

)

// FindAllStudent - Returns total list of registered students.
func (db *DB) FindAllStudent() ([]*model.Student, error) {
	rows, err := db.Query("SELECT * FROM GO_TST.student")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make([]*model.Student, 0)
	for rows.Next() {
		item := new(model.Student)
		err := rows.Scan(&item.Person.ID, &item.Class.ID)
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

// NextIDStudent - Returns the next ID.
func (db *DB) NextIDStudent() ([]*model.Student, error) {
	rows, err := db.Query("SELECT * FROM GO_TST.student")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make([]*model.Student, 0)
	for rows.Next() {
		item := new(model.Student)
		err := rows.Scan(&item.Person.ID, &item.Class.ID)
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

// FindByIDStudent - Returns a specific student by ID.
func (db *DB) FindByIDStudent(id int64) (*model.Student, error) {
	row := db.QueryRow("SELECT * FROM GO_TST.student WHERE person_id=$1", id)
	item := new(model.Student)
	err := row.Scan(&item.Person.ID, &item.Class.ID)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// InsertStudent - Inserts a new student record in the data base.
func (db *DB) InsertStudent(modely model.Student) (int64, error) {
	sqlStatement := "INSERT INTO GO_TST.student (person_id, class_id) VALUES ($1, $2) RETURNING person_id"
	var returnedID int64
	err := db.QueryRow(sqlStatement, modely.Person.ID, modely.Class.ID).Scan(&returnedID)
	if err != nil {
		return 0, err
	}

	return returnedID, nil
}

// UpdateStudent - Updates a base student record.
func (db *DB) UpdateStudent(modely model.Student) error {
	sqlStatement := "UPDATE GO_TST.student SET class_id=$2 WHERE person_id=$1"
	_, err := db.Exec(sqlStatement, modely.Person.ID, modely.Class.ID)
	if err != nil {
		return err
	}
	if err = db.UpdatePerson(modely.Person); err != nil {
		return err
	}
	if err = db.UpdateClass(modely.Class); err != nil {
		return err
	}
	return nil
}
