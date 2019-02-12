package service

import (
	"encoding/json"
	"log"
	"net/http"

	"rest-api/model"
)

// ListarPessoa Retorna lista total pessoas registradas.
func ListarPessoa(w http.ResponseWriter, r *http.Request) {
	db, err := model.NewDB(DataSourcePostgre)
	if err != nil {
		http.Error(w, http.StatusText(405), 405)
		log.Panic(err)
		return
	}

	lista, err := model.ListarPessoa(db)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		log.Panic(err)
		model.CloseDB(db)
		return
	}

	for _, item := range lista {
		log.Println(item.ToString())
	}

	model.CloseDB(db)
	json.NewEncoder(w).Encode(lista)
}
