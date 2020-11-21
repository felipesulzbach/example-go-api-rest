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

func create(entity interface{}, value interface{}) (interface{}, error) {
	db, err := newDB()
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	//query := "INSERT INTO fs_auto.course (name, description, registration_date) VALUES ($1, $2, $3) RETURNING *"
	//row := db.QueryRow(query)

	//"INSERT INTO fs_auto.course (name, description, registration_date) VALUES ($1, $2, $3) RETURNING *"
	val := reflect.ValueOf(entity).Elem()

	var query strings.Builder
	query.WriteString("INSERT INTO fs_auto.")
	query.WriteString(val.Type().Name())
	query.WriteString("(")
	totalFields := val.Type().NumField()
	for i := 0; i < totalFields; i++ {
		//log.Printf("Name: %#v\n", val.Type().Field(i).Name)
		//log.Printf("Type: %#v\n", val.Type().Field(i).Type)
		//log.Printf("Tag: %#v\n", val.Type().Field(i).Tag.Get("db"))
		query.WriteString(val.Type().Field(i).Tag.Get("db"))
		if i < totalFields-1 {
			query.WriteString(",")
		}
	}
	query.WriteString(")")
	query.WriteString(" VALUES (")
	for i := 0; i < totalFields; i++ {
		query.WriteString(getFieldValue(value, val.Type().Field(i)))
		if i < totalFields-1 {
			query.WriteString(",")
		}
	}
	query.WriteString(") RETURNING id")

	log.Println(query.String())

	db.closeDB()
	return "result", nil
}

func getFieldValue(entity interface{}, structField reflect.StructField) string {
	field := getField(entity, structField.Name)

	value := ""
	switch structField.Type {
	case reflect.TypeOf((time.Time)(time.Now())):
		vall := fmt.Sprintf("%v", field)
		dateTime := util.StringToTime(vall)
		if util.TimeIsEmpty(dateTime) {
			value = "null"
		} else {
			value = util.FormatDateTimeISO8601(dateTime)
		}
		break
	case reflect.TypeOf((string)("")):
		value = fmt.Sprintf("%v", field)
		if value == "" {
			value = "null"
		} else {
			value = "'" + value + "'"
		}
	/*case int:
		value, _ = strconv.ParseInt(field.(int64))
		fmt.Println("int:", fieldType)
	case float64:
		fmt.Println("float64:", fieldType)
	case bool:
		fmt.Println("float64:", fieldType)*/
	default:
		value = fmt.Sprintf("%v", field)
		if value == "" {
			value = "null"
		}
	}

	return value
}

func getField(entity interface{}, fieldName string) interface{} {
	reflectValue := reflect.ValueOf(entity)
	value := reflect.Indirect(reflectValue).FieldByName(fieldName)

	return value
}

func update() {
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
