package db

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

// ConnectionTest - Pre-test database connection. If you cannot connect, the server shuts down.
func ConnectionTest() {
	log.Println("DATABASE Testing connection...")
	_, err := _openConnectionDatabase(true)
	if err != nil {
		log.Fatal("SERVER Shutting Down!")
	}
	log.Println("DATABASE Successfully connected!")
}

// Connect ...
func Connect() (*DB, error) {
	db, err := _openConnectionDatabase(false)
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

// Disconnect ...
func (db *DB) Disconnect() {
	defer db.Close()
}

func _openConnectionDatabase(closeTheConnection bool) (*sql.DB, error) {
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
		log.Panicf("DATABASE Error creating connection pool in URI: %s.\n", connDesc.String())
		return nil, err
	}

	if closeTheConnection {
		defer db.Close()
	}

	err = db.Ping()
	if err != nil {
		log.Panicf("DATABASE Error connection in URI: %s.\n", connDesc.String())
		return nil, err
	}
	return db, err
}
