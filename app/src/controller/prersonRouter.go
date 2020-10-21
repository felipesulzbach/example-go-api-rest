package controller

import (
	"github.com/felipesulzbach/exemplo-api-rest/app/src/service"
	"github.com/gorilla/mux"

)

func createRoutersPerson(routerWS *mux.Router) {
	subRouter := routerWS.PathPrefix("/api/v1/person").Subrouter()
	subRouter.Path("").HandlerFunc(service.FindAllPerson).Methods("GET").Name("FindAllPerson")

	printRoutes(subRouter, "Person")
}
