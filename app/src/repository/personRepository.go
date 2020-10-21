package repository

import (
	"log"

	"github.com/felipesulzbach/exemplo-api-rest/app/src/model"

)

// FindAllPerson ...
func FindAllPerson() ([]*model.Person, error) {
	db, err := newDB()
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	rows, err := db.Query("SELECT * FROM fs_auto.person")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make([]*model.Person, 0)
	for rows.Next() {
		item := new(model.Person)
		err := rows.Scan(&item.ID, &item.Name, &item.Cpf, &item.CellPhone, &item.City, &item.ZipCode, &item.Address, &item.RegistrationDate)
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

// FindByIDPerson ...
func FindByIDPerson(id int64) (*model.Person, error) {
	db, err := newDB()
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	row := db.QueryRow("SELECT * FROM fs_auto.person WHERE id=$1", id)

	item := new(model.Person)
	if err := row.Scan(&item.ID, &item.Name, &item.Cpf, &item.CellPhone, &item.City, &item.ZipCode, &item.Address, &item.RegistrationDate); err != nil {
		return nil, err
	}

	db.closeDB()
	return item, nil
}

// InsertPerson ...
func InsertPerson(entity model.Person) (int64, error) {
	db, err := newDB()
	if err != nil {
		log.Panic(err)
		return 0, err
	}

	sqlStatement := "INSERT INTO fs_auto.person (name, cpf, cell_phone, city, zip_code, address, registration_date) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"
	var id int64
	if err := db.QueryRow(sqlStatement, entity.Name, entity.Cpf, entity.CellPhone, entity.City, entity.ZipCode, entity.Address, entity.RegistrationDate).Scan(&id); err != nil {
		return 0, err
	}

	db.closeDB()
	return id, nil
}

// UpdatePerson ...
func UpdatePerson(entity model.Person) error {
	db, err := newDB()
	if err != nil {
		log.Panic(err)
		return err
	}

	sqlStatement := "UPDATE fs_auto.person SET name=$1, cpf=$2, cell_phone=$3, city=$4, zip_code=$5, address=$6, registration_date=$7 WHERE id=$1"
	if _, err := db.Exec(sqlStatement, entity.ID, entity.Name, entity.Cpf, entity.CellPhone, entity.City, entity.ZipCode, entity.Address, entity.RegistrationDate); err != nil {
		return err
	}

	db.closeDB()
	return nil
}
