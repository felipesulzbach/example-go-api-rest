package service

import (
	"net/http"
	"rest-api/model"
	"strconv"

	"github.com/gorilla/mux"
)

// DataSourcePostgre postgres://user:pass@localhost/bookstore?sslmode=disable
//const DataSourcePostgre = "postgres://postgres:admin@localhost/teste_db?sslmode=disable"
const DataSourcePostgre = "postgres://postgres:postgres@localhost/teste_db?sslmode=disable"

//type Env struct{ db model.Datastore }

// Remover Remove um registro da base.
func Remover(w http.ResponseWriter, r *http.Request, db *model.DB, entidade string, coluna string, parametro string) error {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params[parametro], 10, 64)
	if err != nil {
		return err
	}

	if err = model.Remover(db, entidade, coluna, id); err != nil {
		return err
	}
	return nil
}
