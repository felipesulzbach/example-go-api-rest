package service

import (
	"log"
	"net/http"

	"github.com/felipesulzbach/exemplo-api-rest/app/src/repository"

)

// FindAllPerson ...
func FindAllPerson(w http.ResponseWriter, r *http.Request) {
	list, err := repository.FindAllPerson()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		return
	}

	for _, item := range list {
		log.Println(item.ToString())
	}

	jsonOkResponse(w, list)
}
