package model

import (
	"bytes"
	"database/sql"

	_ "github.com/lib/pq"
)

//type Datastore interface {
//	BuscarCursos() ([]*entity.Curso, error)
//}

// DB Conexao db.
type DB struct {
	*sql.DB
}

// NewDB - Opens new PostgreSQL connection.
func NewDB(dataSourceName string) (*DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
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
