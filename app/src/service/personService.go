package service

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/felipesulzbach/exemplo-api-rest/app/src/repository"

)

// FindAllPerson - Returns total list of registered persons.
func FindAllPerson(w http.ResponseWriter, r *http.Request) {
	db, err := repository.NewDB()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		log.Panic(err)
		return
	}

	list, err := db.FindAllPerson()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		db.CloseDB()
		return
	}

	for _, item := range list {
		log.Println(item.ToString())
	}

	db.CloseDB()
	json.NewEncoder(w).Encode(list)
}
