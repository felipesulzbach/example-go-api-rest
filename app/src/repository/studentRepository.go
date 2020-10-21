package repository

import (
	"log"

	"github.com/felipesulzbach/exemplo-api-rest/app/src/model"

)

// FindAllStudent ...
func FindAllStudent() ([]*model.Student, error) {
	db, err := newDB()
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	rows, err := db.Query("SELECT * FROM fs_auto.student")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make([]*model.Student, 0)
	for rows.Next() {
		item := new(model.Student)
		err := rows.Scan(&item.Person.ID, &item.SchoolClass.ID)
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

// FindByIDStudent ...
func FindByIDStudent(id int64) (*model.Student, error) {
	db, err := newDB()
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	row := db.QueryRow("SELECT * FROM fs_auto.student WHERE id=$1", id)
	item := new(model.Student)
	if err := row.Scan(&item.Person.ID, &item.SchoolClass.ID); err != nil {
		return nil, err
	}

	db.closeDB()
	return item, nil
}

// InsertStudent ...
func InsertStudent(entity model.Student) (int64, error) {
	db, err := newDB()
	if err != nil {
		log.Panic(err)
		return 0, err
	}

	sqlStatement := "INSERT INTO fs_auto.student (school_class_id) VALUES ($1) RETURNING id"
	var id int64
	if err := db.QueryRow(sqlStatement, entity.SchoolClass.ID).Scan(&id); err != nil {
		return 0, err
	}

	db.closeDB()
	return id, nil
}

// UpdateStudent ...
func UpdateStudent(entity model.Student) error {
	db, err := newDB()
	if err != nil {
		log.Panic(err)
		return err
	}

	sqlStatement := "UPDATE fs_auto.student SET school_class_id=$1 WHERE id=$1"
	if _, err := db.Exec(sqlStatement, entity.Person.ID, entity.SchoolClass.ID); err != nil {
		return err
	}

	db.closeDB()
	return nil
}
