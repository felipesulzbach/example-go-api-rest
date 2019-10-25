package model

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5435
	user     = "postgres"
	password = "postgres"
	dbname   = "go_rest_db"
	driver   = "postgres"
)

//type Datastore interface {
//	BuscarCursos() ([]*entity.Curso, error)
//}

// DB Database connection.
type DB struct {
	*sql.DB
}

// TestConnectionDB - Makes a database connection test. If you cannot connect, the server shuts down.
func TestConnectionDB() {
	connString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var connDesc strings.Builder
	connDesc.WriteString(host)
	connDesc.WriteString(":")
	connDesc.WriteString(strconv.FormatInt(port, 10))
	connDesc.WriteString("/")
	connDesc.WriteString(dbname)

	log.Println("DATABASE Testing connection...")

	db, err := sql.Open(driver, connString)
	if err != nil {
		log.Fatalf("DATABASE Error creating connection pool in URI: %s. The connection settings are in file '.../model/model.go'.\n", connDesc.String())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("DATABASE Error connection in URI: %s. The connection settings are in file '.../model/model.go'.\n", connDesc.String())
	}

	log.Println("DATABASE Successfully connected!")
}

// NewDB - Opens new PostgreSQL connection.
func NewDB() (*DB, error) {
	connString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open(driver, connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
		return nil, err
	}
	if err = db.Ping(); err != nil {
		log.Fatal("Error connection in Database: ", err.Error())
		return nil, err
	}
	return &DB{db}, nil
}

// CloseDB - Closes the connection to PostgreSQL.
func CloseDB(db *DB) {
	defer db.Close()
}

// Delete - Removes a record from the base.
func Delete(db *DB, entity string, column string, value int64) error {
	var sqlStatement bytes.Buffer
	sqlStatement.WriteString("DELETE FROM GO_TST.")
	sqlStatement.WriteString(entity)
	sqlStatement.WriteString(" WHERE ")
	sqlStatement.WriteString(column)
	sqlStatement.WriteString("=$1")
	_, err := db.Exec(sqlStatement.String(), value)
	if err != nil {
		return err
	}
	return nil
}
