package controller

import (
	"github.com/felipesulzbach/exemplo-api-rest/app/src/service"
	"github.com/gorilla/mux"

)

func createRoutersCourse(routerWS *mux.Router) {
	subRouter := routerWS.PathPrefix("/api/v1/course").Subrouter()
	subRouter.Path("").HandlerFunc(service.FindAllCourse).Methods("GET").Name("FindAllCourse")
	subRouter.Path("/{id}").HandlerFunc(service.FindByIDCourse).Methods("GET").Name("FindByIDCourse")
	subRouter.Path("").HandlerFunc(service.InsertCourse).Methods("POST").Name("InsertCourse")
	subRouter.Path("").HandlerFunc(service.UpdateCourse).Methods("PUT").Name("UpdateCourse")
	subRouter.Path("/{id}").HandlerFunc(service.DeleteCourse).Methods("DELETE").Name("DeleteCourse")

	printRoutes(subRouter, "Course")
}
