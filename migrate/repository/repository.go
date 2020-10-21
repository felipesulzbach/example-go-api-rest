package repository

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

	_ "github.com/lib/pq"

)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "go_rest_db"
	driver   = "postgres"
)

// DB Database connection.
type DB struct {
	*sql.DB
}

// OpenDB ...
func OpenDB() (*DB, error) {
	db, err := openConnectionDatabase(false)
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

func openConnectionDatabase(closeTheConnection bool) (*sql.DB, error) {
	connString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var connDesc strings.Builder
	connDesc.WriteString(host)
	connDesc.WriteString(":")
	connDesc.WriteString(strconv.FormatInt(port, 10))
	connDesc.WriteString("/")
	connDesc.WriteString(dbname)

	db, err := sql.Open(driver, connString)
	if err != nil {
		log.Panicf("DATABASE Error creating connection pool in URI: %s. The connection settings are in file '.../migrate/repository.go'.\n", connDesc.String())
		return nil, err
	}

	if closeTheConnection {
		defer db.Close()
	}

	err = db.Ping()
	if err != nil {
		log.Panicf("DATABASE Error connection in URI: %s. The connection settings are in file '.../migrate/repository.go'.\n", connDesc.String())
		return nil, err
	}
	return db, err
}

// CloseDB ...
func (db *DB) CloseDB() {
	defer db.Close()
}
