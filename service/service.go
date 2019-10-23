package service

import (
	"net/http"
	"strconv"

	"github.com/_dev/exemplo-api-rest/model"

	"github.com/gorilla/mux"
)

// DataSourcePostgre - URI of the database.
const DataSourcePostgre = "postgres://postgres:postgres@localhost/go_rest_db?sslmode=disable"

//type Env struct{ db model.Datastore }

// Delete - Removes a record from the base.
func Delete(w http.ResponseWriter, r *http.Request, db *model.DB, entity string, column string, parametter string) error {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params[parametter], 10, 64)
	if err != nil {
		return err
	}

	if err = model.Delete(db, entity, column, id); err != nil {
		return err
	}
	return nil
}
