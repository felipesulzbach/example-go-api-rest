package controller

import (
	"log"
	"net/http"

	"github.com/felipesulzbach/exemplo-api-rest/app/src/service"

)

func getAllPerson(w http.ResponseWriter, r *http.Request) {
	response, err := service.FindAllPerson()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		return
	}

	jsonOkResponse(w, response)
}
