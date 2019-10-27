package model

import (
	"github.com/_dev/exemplo-api-rest/model/entity"
)

// NextIDPerson - Returns the next ID.
func (db *DB) NextIDPerson() (int64, error) {
	row := db.QueryRow("SELECT (MAX(id) + 1) FROM GO_TST.person")

	var id int64
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// FindAllPerson - Returns total list of registered persons.
func (db *DB) FindAllPerson() ([]*entity.Person, error) {
	rows, err := db.Query("SELECT * FROM GO_TST.person")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make([]*entity.Person, 0)
	for rows.Next() {
		item := new(entity.Person)
		err := rows.Scan(&item.ID, &item.Name, &item.Cpf, &item.CellPhone, &item.City, &item.ZipCode, &item.Address, &item.RegistrationDate)
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

// FindByIDPerson - Returns a specific person by ID.
func (db *DB) FindByIDPerson(id int64) (*entity.Person, error) {
	row := db.QueryRow("SELECT * FROM GO_TST.person WHERE id=$1", id)

	item := new(entity.Person)
	err := row.Scan(&item.ID, &item.Name, &item.Cpf, &item.CellPhone, &item.City, &item.ZipCode, &item.Address, &item.RegistrationDate)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// InsertPerson - Inserts a new person record in the data base.
func (db *DB) InsertPerson(entityy entity.Person) (int64, error) {
	sqlStatement := "INSERT INTO GO_TST.person (id, name, cpf, cell_phone, city, zip_code, address, registration_date) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id"
	var returnedID int64
	err := db.QueryRow(sqlStatement, entityy.ID, entityy.Name, entityy.Cpf, entityy.CellPhone, entityy.City, entityy.ZipCode, entityy.Address, entityy.RegistrationDate).Scan(&returnedID)
	if err != nil {
		return 0, err
	}

	return returnedID, nil
}

// UpdatePerson - Updates a base class record.
func (db *DB) UpdatePerson(entityy entity.Person) error {
	sqlStatement := "UPDATE GO_TST.person SET name=$2, cpf=$3, cell_phone=$4, city=$5, zip_code=$6, address=$7, registration_date=$8 WHERE id=$1"
	_, err := db.Exec(sqlStatement, entityy.ID, entityy.Name, entityy.Cpf, entityy.CellPhone, entityy.City, entityy.ZipCode, entityy.Address, entityy.RegistrationDate)
	if err != nil {
		return err
	}
	return nil
}
