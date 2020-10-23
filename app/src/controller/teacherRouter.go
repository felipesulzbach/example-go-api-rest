package controller

import (
	"github.com/gorilla/mux"

)

func createRoutersTeacher(routerWS *mux.Router) {
	subRouter := routerWS.PathPrefix("/api/v1/teacher").Subrouter()
	subRouter.Path("").HandlerFunc(getAllTeacher).Methods("GET").Name("FindAllTeacher")
	subRouter.Path("/{id}").HandlerFunc(getByIDTeacher).Methods("GET").Name("FindByIDTeacher")
	subRouter.Path("").HandlerFunc(insertTeacher).Methods("POST").Name("InsertTeacher")
	subRouter.Path("").HandlerFunc(updateTeacher).Methods("PUT").Name("UpdateTeacher")
	subRouter.Path("/{id}").HandlerFunc(deleteTeacher).Methods("DELETE").Name("DeleteTeacher")

	printRoutes(subRouter, "Teacher")
}
