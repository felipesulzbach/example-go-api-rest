package controller

import (
	"github.com/gorilla/mux"

)

func createRoutersPerson(routerWS *mux.Router) {
	subRouter := routerWS.PathPrefix("/api/v1/person").Subrouter()
	subRouter.Path("").HandlerFunc(getAllPerson).Methods("GET").Name("FindAllPerson")

	printRoutes(subRouter, "Person")
}
