package controller

import (
	"log"
	"net/http"

	"github.com/_dev/exemplo-api-rest/service"

	"github.com/gorilla/mux"
)

const (
	host   = "localhost"
	port   = "8080"
	Domain = host + ":" + port
)

// CreateRouters - Create routes to endpoints.
func CreateRouters(routerWS *mux.Router) {
	log.Println("ROUTERS Creating...")

	createRoutersCourse(routerWS)

	subRouterCourse := routerWS.PathPrefix("/course").Subrouter()
	subRouterCourse.Path("").HandlerFunc(service.FindAllCourse).Methods("GET").Name("FindAllCourse")
	subRouterCourse.Path("/{id}").HandlerFunc(service.FindByIDCourse).Methods("GET").Name("FindByIDCourse")
	subRouterCourse.Path("").HandlerFunc(service.InsertCourse).Methods("POST").Name("InsertCourse")
	subRouterCourse.Path("").HandlerFunc(service.UpdateCourse).Methods("PUT").Name("UpdateCourse")
	subRouterCourse.Path("/{id}").HandlerFunc(service.DeleteCourse).Methods("DELETE").Name("DeleteCourse")

	subRouterClass := routerWS.PathPrefix("/class").Subrouter()
	subRouterClass.Path("").HandlerFunc(service.FindAllClass).Methods("GET").Name("FindAllClass")
	subRouterClass.Path("/{id}").HandlerFunc(service.FindByIDClass).Methods("GET").Name("FindByIDClass")
	subRouterClass.Path("").HandlerFunc(service.InsertClass).Methods("POST").Name("InsertClass")
	subRouterClass.Path("").HandlerFunc(service.UpdateClass).Methods("PUT").Name("UpdateClass")
	subRouterClass.Path("/{id}").HandlerFunc(service.DeleteClass).Methods("DELETE").Name("DeleteClass")

	subRouterPerson := routerWS.PathPrefix("/person").Subrouter()
	subRouterPerson.Path("").HandlerFunc(service.FindAllPerson).Methods("GET").Name("FindAllPerson")

	subRouterStudent := routerWS.PathPrefix("/student").Subrouter()
	subRouterStudent.Path("").HandlerFunc(service.FindAllStudent).Methods("GET").Name("FindAllStudent")
	subRouterStudent.Path("/{id}").HandlerFunc(service.FindByIDStudent).Methods("GET").Name("FindByIDStudent")
	subRouterStudent.Path("").HandlerFunc(service.InsertStudent).Methods("POST").Name("InsertStudent")
	subRouterStudent.Path("").HandlerFunc(service.UpdateStudent).Methods("PUT").Name("UpdateStudent")
	subRouterStudent.Path("/{id}").HandlerFunc(service.DeleteStudent).Methods("DELETE").Name("DeleteStudent")

	subRouterTeacher := routerWS.PathPrefix("/teacher").Subrouter()
	subRouterTeacher.Path("").HandlerFunc(service.FindAllTeacher).Methods("GET").Name("FindAllTeacher")
	subRouterTeacher.Path("/{id}").HandlerFunc(service.FindByIDTeacher).Methods("GET").Name("FindByIDTeacher")
	subRouterTeacher.Path("").HandlerFunc(service.InsertTeacher).Methods("POST").Name("InsertTeacher")
	subRouterTeacher.Path("").HandlerFunc(service.UpdateTeacher).Methods("PUT").Name("UpdateTeacher")
	subRouterTeacher.Path("/{id}").HandlerFunc(service.DeleteTeacher).Methods("DELETE").Name("DeleteTeacher")

	http.Handle("/", routerWS)

	log.Printf("ROUTERS Successfully created on: http://%s/", Domain)
}

// CreateRoutersCourse - Creates routes to Course endpoints.
func createRoutersCourse(routerWS *mux.Router) {
	subRouterCourse := routerWS.PathPrefix("/course").Subrouter()
	subRouterCourse.Path("").HandlerFunc(service.FindAllCourse).Methods("GET").Name("FindAllCourse")
	subRouterCourse.Path("/{id}").HandlerFunc(service.FindByIDCourse).Methods("GET").Name("FindByIDCourse")
	subRouterCourse.Path("").HandlerFunc(service.InsertCourse).Methods("POST").Name("InsertCourse")
	subRouterCourse.Path("").HandlerFunc(service.UpdateCourse).Methods("PUT").Name("UpdateCourse")
	subRouterCourse.Path("/{id}").HandlerFunc(service.DeleteCourse).Methods("DELETE").Name("DeleteCourse")
}
