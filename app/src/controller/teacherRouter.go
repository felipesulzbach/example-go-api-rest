package controller

import (
	"github.com/felipesulzbach/exemplo-api-rest/app/src/service"
	"github.com/gorilla/mux"

)

func createRoutersTeacher(routerWS *mux.Router) {
	subRouter := routerWS.PathPrefix("/api/v1/teacher").Subrouter()
	subRouter.Path("").HandlerFunc(service.FindAllTeacher).Methods("GET").Name("FindAllTeacher")
	subRouter.Path("/{id}").HandlerFunc(service.FindByIDTeacher).Methods("GET").Name("FindByIDTeacher")
	subRouter.Path("").HandlerFunc(service.InsertTeacher).Methods("POST").Name("InsertTeacher")
	subRouter.Path("").HandlerFunc(service.UpdateTeacher).Methods("PUT").Name("UpdateTeacher")
	subRouter.Path("/{id}").HandlerFunc(service.DeleteTeacher).Methods("DELETE").Name("DeleteTeacher")

	printRoutes(subRouter, "Teacher")
}
