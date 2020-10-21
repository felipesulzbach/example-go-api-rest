package controller

import (
	"github.com/felipesulzbach/exemplo-api-rest/app/src/service"
	"github.com/gorilla/mux"

)

func createRoutersSchoolClass(routerWS *mux.Router) {
	subRouter := routerWS.PathPrefix("/api/v1/schoolclass").Subrouter()
	subRouter.Path("").HandlerFunc(service.FindAllSchoolClass).Methods("GET").Name("FindAllSchoolClass")
	subRouter.Path("/{id}").HandlerFunc(service.FindByIDSchoolClass).Methods("GET").Name("FindByIDSchoolClass")
	subRouter.Path("").HandlerFunc(service.InsertSchoolClass).Methods("POST").Name("InsertSchoolClass")
	subRouter.Path("").HandlerFunc(service.UpdateSchoolClass).Methods("PUT").Name("UpdateSchoolClass")
	subRouter.Path("/{id}").HandlerFunc(service.DeleteSchoolClass).Methods("DELETE").Name("DeleteSchoolClass")

	printRoutes(subRouter, "SchoolClass")
}
