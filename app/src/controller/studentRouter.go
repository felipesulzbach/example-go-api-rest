package controller

import (
	"github.com/gorilla/mux"

)

func createRoutersStudent(routerWS *mux.Router) {
	subRouter := routerWS.PathPrefix("/api/v1/student").Subrouter()
	subRouter.Path("").HandlerFunc(getAllStudent).Methods("GET").Name("FindAllStudent")
	subRouter.Path("/{id}").HandlerFunc(getByIDStudent).Methods("GET").Name("FindByIDStudent")
	subRouter.Path("").HandlerFunc(createStudent).Methods("POST").Name("CreateStudent")
	subRouter.Path("").HandlerFunc(updateStudent).Methods("PUT").Name("UpdateStudent")
	subRouter.Path("/{id}").HandlerFunc(deleteStudent).Methods("DELETE").Name("DeleteStudent")

	printRoutes(subRouter, "Student")
}
