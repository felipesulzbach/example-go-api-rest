package controller

import (
	"github.com/gorilla/mux"

)

func createRoutersSchoolClass(routerWS *mux.Router) {
	subRouter := routerWS.PathPrefix("/api/v1/schoolclass").Subrouter()
	subRouter.Path("").HandlerFunc(getAllSchoolClass).Methods("GET").Name("FindAllSchoolClass")
	subRouter.Path("/{id}").HandlerFunc(getByIDSchoolClass).Methods("GET").Name("FindByIDSchoolClass")
	subRouter.Path("").HandlerFunc(insertSchoolClass).Methods("POST").Name("InsertSchoolClass")
	subRouter.Path("").HandlerFunc(updateSchoolClass).Methods("PUT").Name("UpdateSchoolClass")
	subRouter.Path("/{id}").HandlerFunc(deleteSchoolClass).Methods("DELETE").Name("DeleteSchoolClass")

	printRoutes(subRouter, "SchoolClass")
}
