package model

import (
	"github.com/_dev/exemplo-api-rest/model/entity"
)

// FindAllStudent - Returns total list of registered students.
func (db *DB) FindAllStudent() ([]*entity.Student, error) {
	rows, err := db.Query("SELECT * FROM GO_TST.student")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make([]*entity.Student, 0)
	for rows.Next() {
		item := new(entity.Student)
		err := rows.Scan(&item.Person, &item.Class)
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
func (db *DB) NextIDStudent() ([]*entity.Student, error) {
	rows, err := db.Query("SELECT * FROM GO_TST.student")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make([]*entity.Student, 0)
	for rows.Next() {
		item := new(entity.Student)
		err := rows.Scan(&item.Person, &item.Class)
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
func (db *DB) FindByIDStudent(id int64) (*entity.Student, error) {
	row := db.QueryRow("SELECT * FROM GO_TST.student WHERE person_id=$1", id)
	item := new(entity.Student)
	err := row.Scan(&item.Person, &item.Class)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// InsertStudent - Inserts a new student record in the data base.
func (db *DB) InsertStudent(entityy entity.Student) (int64, error) {
	sqlStatement := "INSERT INTO GO_TST.student (person_id, class_id) VALUES ($1, $2) RETURNING person_id"
	var returnedID int64
	err := db.QueryRow(sqlStatement, entityy.Person.ID, entityy.Class.ID).Scan(&returnedID)
	if err != nil {
		return 0, err
	}

	return returnedID, nil
}

// UpdateStudent - Updates a base student record.
func (db *DB) UpdateStudent(entityy entity.Student) error {
	sqlStatement := "UPDATE GO_TST.student SET class_id=$2 WHERE person_id=$1"
	_, err := db.Exec(sqlStatement, entityy.Person.ID, entityy.Class.ID)
	if err != nil {
		return err
	}
	if err = db.UpdatePerson(entityy.Person); err != nil {
		return err
	}
	if err = db.UpdateClass(entityy.Class); err != nil {
		return err
	}
	return nil
}
