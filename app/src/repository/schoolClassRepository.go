package repository

import (
	"log"

	"github.com/felipesulzbach/exemplo-api-rest/app/src/model"

)

// FindAllSchoolClass ...
func FindAllSchoolClass() ([]*model.SchoolClass, error) {
	db, err := newDB()
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	rows, err := db.Query("SELECT * FROM fs_auto.school_class")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make([]*model.SchoolClass, 0)
	for rows.Next() {
		item := new(model.SchoolClass)
		err := rows.Scan(&item.ID, &item.Course.ID, &item.StartDate, &item.EndDate, &item.RegistrationDate)
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

// FindByIDSchoolClass ...
func FindByIDSchoolClass(id int64) (*model.SchoolClass, error) {
	db, err := newDB()
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	row := db.QueryRow("SELECT * FROM fs_auto.school_class WHERE id=$1", id)

	item := new(model.SchoolClass)
	if err := row.Scan(&item.ID, &item.Course.ID, &item.StartDate, &item.EndDate, &item.RegistrationDate); err != nil {
		return nil, err
	}

	db.closeDB()
	return item, nil
}

// InsertSchoolClass ...
func InsertSchoolClass(modell model.SchoolClass) (int64, error) {
	db, err := newDB()
	if err != nil {
		log.Panic(err)
		return 0, err
	}

	sqlStatement := "INSERT INTO fs_auto.school_class (course_id, start_date, end_date, registration_date) VALUES ($1, $2, $3, $4) RETURNING id"
	var returnedID int64
	if err := db.QueryRow(sqlStatement, modell.Course.ID, modell.StartDate, modell.EndDate, modell.RegistrationDate).Scan(&returnedID); err != nil {
		return 0, err
	}

	db.closeDB()
	return returnedID, nil
}

// UpdateSchoolClass ...
func UpdateSchoolClass(entity model.SchoolClass) error {
	db, err := newDB()
	if err != nil {
		log.Panic(err)
		return err
	}

	sqlStatement := "UPDATE fs_auto.school_class SET course_id=$1, start_date=$2, end_date=$3, registration_date=$4 WHERE id=$1"
	if _, err := db.Exec(sqlStatement, entity.ID, entity.Course.ID, entity.StartDate, entity.EndDate, entity.RegistrationDate); err != nil {
		return err
	}

	db.closeDB()
	return nil
}
