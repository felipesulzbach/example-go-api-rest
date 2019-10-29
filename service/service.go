package service

import (
	"net/http"
	"strconv"

	"github.com/_dev/exemplo-api-rest/model"

	"github.com/gorilla/mux"
)

// ServiceDelete - Parameters usage in delete.
type ServiceDelete struct {
	Entity     string
	Column     string
	Parametter string
}

// Delete - Removes a record from the base.
func (serviceDelete ServiceDelete) Delete(w http.ResponseWriter, r *http.Request, db *model.DB) error {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params[serviceDelete.Parametter], 10, 64)
	if err != nil {
		return err
	}

	if err = db.Delete(serviceDelete.Entity, serviceDelete.Column, id); err != nil {
		return err
	}
	return nil
}
