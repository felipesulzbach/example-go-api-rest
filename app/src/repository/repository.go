package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/felipesulzbach/exemplo-api-rest/app/src/util"

)

func getByID(tableName string, id int64) (interface{}, error) {
	objectMap, err := getAllByArgs("SELECT * FROM fs_auto."+tableName+" WHERE id=$1", id)
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	if len(objectMap) > 1 {
		return nil, errors.New("The query returned more than one result")
	}

	var result interface{}
	for _, object := range objectMap {
		result = object
	}

	return result, nil
}

func getAll(tableName string) ([]map[string]interface{}, error) {
	query := "SELECT * FROM fs_auto." + tableName
	objectMap, err := getAllByArgs(query, nil)
	if err != nil {
		return nil, err
	}

	return objectMap, nil
}

func getAllByArgs(query string, args interface{}) ([]map[string]interface{}, error) {
	db, err := newDB()
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	var rows *sql.Rows
	if args != nil {
		rows, err = db.Query(query, args)
	} else {
		rows, err = db.Query(query)
	}
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

func create(entity interface{}, value interface{}) (int64, error) {
	db, err := newDB()
	if err != nil {
		log.Panic(err)
		return 0, err
	}

	val := reflect.ValueOf(entity).Elem()
	columns, values := _getQueryFields(val, value)

	var query strings.Builder
	query.WriteString("INSERT INTO fs_auto.")
	query.WriteString(val.Type().Name())
	query.WriteString("(")
	query.WriteString(columns.String())
	query.WriteString(") VALUES (")
	query.WriteString(values.String())
	query.WriteString(") RETURNING id")

	var id int64
	if err := db.QueryRow(query.String()).Scan(&id); err != nil {
		return 0, err
	}

	db.closeDB()
	return id, nil
}

func update(entity interface{}) error {
	query := _getQueryUpdate(entity)

	db, err := newDB()
	if err != nil {
		log.Panic(err)
		return err
	}

	if _, err := db.Exec(query); err != nil {
		return err
	}

	db.closeDB()
	return nil
}

func delete(tableName string, id int64) error {
	db, err := newDB()
	if err != nil {
		log.Panic(err)
		return err
	}

	query := "DELETE fs_auto." + tableName + " WHERE id=$1"
	if _, err := db.Exec(query, id); err != nil {
		return err
	}

	db.closeDB()
	return nil
}

func _getQueryFields(val reflect.Value, value interface{}) (strings.Builder, strings.Builder) {
	var columns strings.Builder
	var values strings.Builder

	totalFields := val.Type().NumField()
	for i := 0; i < totalFields; i++ {
		if val.Type().Field(i).Tag.Get("db") == "id" && _getFieldValue(value, val.Type().Field(i)) == "0" {
			continue
		}

		columns.WriteString(val.Type().Field(i).Tag.Get("db"))
		values.WriteString(_getFieldValue(value, val.Type().Field(i)))
		if i < totalFields-1 {
			columns.WriteString(",")
			values.WriteString(",")
		}
	}

	return columns, values
}

func _getFieldValue(entity interface{}, structField reflect.StructField) string {
	field := _getField(entity, structField.Name)

	value := ""
	switch structField.Type {
	case reflect.TypeOf((time.Time)(time.Now())):
		vall := fmt.Sprintf("%v", field)
		dateTime := util.StringToTime(vall)
		if util.IsDateTimeEmpty(dateTime) {
			value = "null"
		} else {
			value = "'" + util.FormatDateTimeISO8601(dateTime) + "'"
		}
		break
	case reflect.TypeOf((string)("")):
		value = fmt.Sprintf("%v", field)
		if value == "" {
			value = "null"
		} else {
			value = "'" + value + "'"
		}
	default:
		value = fmt.Sprintf("%v", field)
		if value == "" {
			value = "null"
		}
	}

	return value
}

func _getField(entity interface{}, fieldName string) interface{} {
	reflectValue := reflect.ValueOf(entity)
	value := reflect.Indirect(reflectValue).FieldByName(fieldName)

	return value
}

func _getQueryUpdate(entity interface{}) string {
	val := reflect.ValueOf(entity).Elem()

	var query strings.Builder
	query.WriteString("UPDATE fs_auto.")
	query.WriteString(val.Type().Name())
	query.WriteString(" SET ")

	totalFields := val.Type().NumField()
	var id string
	for i := 0; i < totalFields; i++ {
		if val.Type().Field(i).Tag.Get("db") == "id" {
			id = _getFieldValue(entity, val.Type().Field(i))
			continue
		}

		query.WriteString(val.Type().Field(i).Tag.Get("db"))
		query.WriteString("=")
		query.WriteString(_getFieldValue(entity, val.Type().Field(i)))
		if i < totalFields-1 {
			query.WriteString(",")
		}
	}
	query.WriteString(" WHERE id=")
	query.WriteString(id)

	return query.String()
}
