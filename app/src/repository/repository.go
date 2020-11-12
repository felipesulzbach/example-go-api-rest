package repository

import (
	"log"
	"reflect"

)

/*
func getByID(id int64, entity interface{}) (interface{}, error) {
	db, err := newDB()
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	row := db.QueryRow("SELECT * FROM fs_auto."+entity.GetTableName()+" WHERE id=$1", id)

	result := new(entity)
	if err := row.Scan(&result.ID, &result.Name, &result.Description, &result.RegistrationDate); err != nil {
		return nil, err
	}

	db.closeDB()
	return result, nil
}
*/
func getAll(query string) ([]map[string]interface{}, error) {
	db, err := newDB()
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]map[string]interface{}, 0)
	for rows.Next() {
		columns, err := rows.ColumnTypes()
		if err != nil {
			log.Fatalln(err)
		}

		values := make([]interface{}, len(columns))
		object := map[string]interface{}{}
		for i, column := range columns {
			object[column.Name()] = reflect.New(column.ScanType()).Interface()
			values[i] = object[column.Name()]
		}

		err = rows.Scan(values...)
		if err != nil {
			return nil, err
		}

		result = append(result, object)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	db.closeDB()
	return result, nil
}

func getOne() {
}

func save() {
}

func delete() {
}
