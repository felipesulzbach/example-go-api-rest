package controller

import (
	"github.com/felipesulzbach/exemplo-api-rest/app/src/service"
	"github.com/gorilla/mux"

)

func createRoutersStudent(routerWS *mux.Router) {
	subRouter := routerWS.PathPrefix("/api/v1/student").Subrouter()
	subRouter.Path("").HandlerFunc(service.FindAllStudent).Methods("GET").Name("FindAllStudent")
	subRouter.Path("/{id}").HandlerFunc(service.FindByIDStudent).Methods("GET").Name("FindByIDStudent")
	subRouter.Path("").HandlerFunc(service.InsertStudent).Methods("POST").Name("InsertStudent")
	subRouter.Path("").HandlerFunc(service.UpdateStudent).Methods("PUT").Name("UpdateStudent")
	subRouter.Path("/{id}").HandlerFunc(service.DeleteStudent).Methods("DELETE").Name("DeleteStudent")

	printRoutes(subRouter, "Student")
}
