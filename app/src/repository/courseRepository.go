package repository

import (
	"log"

	"github.com/felipesulzbach/exemplo-api-rest/app/src/model"

)

// FindAllCourse ...
func FindAllCourse() ([]*model.Course, error) {
	db, err := newDB()
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	rows, err := db.Query("SELECT * FROM fs_auto.course")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make([]*model.Course, 0)
	for rows.Next() {
		item := new(model.Course)
		err := rows.Scan(&item.ID, &item.Name, &item.Description, &item.RegistrationDate)
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

// FindByIDCourse ...
func FindByIDCourse(id int64) (*model.Course, error) {
	db, err := newDB()
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	row := db.QueryRow("SELECT * FROM fs_auto.course WHERE id=$1", id)

	item := new(model.Course)
	if err := row.Scan(&item.ID, &item.Name, &item.Description, &item.RegistrationDate); err != nil {
		return nil, err
	}

	db.closeDB()
	return item, nil
}

// InsertCourse ...
func InsertCourse(entity model.Course) (int64, error) {
	db, err := newDB()
	if err != nil {
		log.Panic(err)
		return 0, err
	}

	sqlStatement := "INSERT INTO fs_auto.course (name, description, registration_date) VALUES ($1, $2, $3) RETURNING id"
	var id int64
	if err := db.QueryRow(sqlStatement, entity.Name, entity.Description, entity.RegistrationDate).Scan(&id); err != nil {
		return 0, err
	}

	db.closeDB()
	return id, nil
}

// UpdateCourse ...
func UpdateCourse(entity model.Course) error {
	db, err := newDB()
	if err != nil {
		log.Panic(err)
		return err
	}

	sqlStatement := "UPDATE fs_auto.course SET name=$1, description=$2, registration_date=$3 WHERE id=$1"
	if _, err := db.Exec(sqlStatement, entity.ID, entity.Name, entity.Description, entity.RegistrationDate); err != nil {
		return err
	}

	db.closeDB()
	return nil
}
