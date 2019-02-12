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

// NewDB Abre nova conexao com o PostgreSQL.
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

// CloseDB Fecha a conexao com o PostgreSQL.
func CloseDB(db *DB) {
	defer db.Close()
}

// Remover Remove um registro da base.
func Remover(db *DB, entity string, coluna string, valor int64) error {
	var sqlStatement bytes.Buffer
	sqlStatement.WriteString("DELETE FROM ")
	sqlStatement.WriteString(entity)
	sqlStatement.WriteString(" WHERE ")
	sqlStatement.WriteString(coluna)
	sqlStatement.WriteString("=$1")
	_, err := db.Exec(sqlStatement.String(), valor)
	if err != nil {
		return err
	}
	return nil
}
