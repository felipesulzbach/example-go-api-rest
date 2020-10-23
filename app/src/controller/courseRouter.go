package controller

import (
	"github.com/gorilla/mux"

)

func createRoutersCourse(routerWS *mux.Router) {
	subRouter := routerWS.PathPrefix("/api/v1/course").Subrouter()
	subRouter.Path("").HandlerFunc(getAllCourse).Methods("GET").Name("FindAllCourse")
	subRouter.Path("/{id}").HandlerFunc(getByIDCourse).Methods("GET").Name("FindByIDCourse")
	subRouter.Path("").HandlerFunc(createCourse).Methods("POST").Name("InsertCourse")
	subRouter.Path("").HandlerFunc(updateCourse).Methods("PUT").Name("UpdateCourse")
	subRouter.Path("/{id}").HandlerFunc(deleteCourse).Methods("DELETE").Name("DeleteCourse")

	printRoutes(subRouter, "Course")
}
