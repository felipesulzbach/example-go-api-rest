package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/felipesulzbach/exemplo-api-rest/app/src/repository"
	"github.com/gorilla/mux"

)

// DeletePO - Parameters usage in delete.
type deletePO struct {
	entity     string
	column     string
	parametter string
}

// Delete - Removes a record from the base.
func (po deletePO) Delete(w http.ResponseWriter, r *http.Request, db *repository.DB) error {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params[po.parametter], 10, 64)
	if err != nil {
		return err
	}

	if err = db.Delete(po.entity, po.column, id); err != nil {
		return err
	}
	return nil
}

func jsonOkResponse(w http.ResponseWriter, value interface{}) {
	jsonResponse(w, value, 200)
}

func jsonCreatedResponse(w http.ResponseWriter, value interface{}) {
	jsonResponse(w, value, 201)
}

func jsonResponse(w http.ResponseWriter, value interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(value)
}
