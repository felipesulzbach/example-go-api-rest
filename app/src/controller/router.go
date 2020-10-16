package controller

import (
	"log"
	"net/http"

	"github.com/felipesulzbach/exemplo-api-rest/app/src/service"
	"github.com/gorilla/mux"

)

const (
	host = "localhost"
	port = "8080"

	// Domain - Domain to connect to the database.
	Domain = host + ":" + port
)

// CreateRouters - Create routes to endpoints.
func CreateRouters(routerWS *mux.Router) {
	log.Println("ROUTERS Creating...")

	createRoutersCourse(routerWS)
	createRoutersClass(routerWS)
	createRoutersPerson(routerWS)
	createRoutersStudent(routerWS)
	createRoutersTeacher(routerWS)
	http.Handle("/", routerWS)

	log.Printf("ROUTERS Successfully created on: http://%s/", Domain)
}

func createRoutersCourse(routerWS *mux.Router) {
	subRouter := routerWS.PathPrefix("/course").Subrouter()
	subRouter.Path("").HandlerFunc(service.FindAllCourse).Methods("GET").Name("FindAllCourse")
	subRouter.Path("/{id}").HandlerFunc(service.FindByIDCourse).Methods("GET").Name("FindByIDCourse")
	subRouter.Path("").HandlerFunc(service.InsertCourse).Methods("POST").Name("InsertCourse")
	subRouter.Path("").HandlerFunc(service.UpdateCourse).Methods("PUT").Name("UpdateCourse")
	subRouter.Path("/{id}").HandlerFunc(service.DeleteCourse).Methods("DELETE").Name("DeleteCourse")
}

func createRoutersClass(routerWS *mux.Router) {
	subRouter := routerWS.PathPrefix("/class").Subrouter()
	subRouter.Path("").HandlerFunc(service.FindAllClass).Methods("GET").Name("FindAllClass")
	subRouter.Path("/{id}").HandlerFunc(service.FindByIDClass).Methods("GET").Name("FindByIDClass")
	subRouter.Path("").HandlerFunc(service.InsertClass).Methods("POST").Name("InsertClass")
	subRouter.Path("").HandlerFunc(service.UpdateClass).Methods("PUT").Name("UpdateClass")
	subRouter.Path("/{id}").HandlerFunc(service.DeleteClass).Methods("DELETE").Name("DeleteClass")
}

func createRoutersPerson(routerWS *mux.Router) {
	subRouter := routerWS.PathPrefix("/person").Subrouter()
	subRouter.Path("").HandlerFunc(service.FindAllPerson).Methods("GET").Name("FindAllPerson")
}

func createRoutersStudent(routerWS *mux.Router) {
	subRouter := routerWS.PathPrefix("/student").Subrouter()
	subRouter.Path("").HandlerFunc(service.FindAllStudent).Methods("GET").Name("FindAllStudent")
	subRouter.Path("/{id}").HandlerFunc(service.FindByIDStudent).Methods("GET").Name("FindByIDStudent")
	subRouter.Path("").HandlerFunc(service.InsertStudent).Methods("POST").Name("InsertStudent")
	subRouter.Path("").HandlerFunc(service.UpdateStudent).Methods("PUT").Name("UpdateStudent")
	subRouter.Path("/{id}").HandlerFunc(service.DeleteStudent).Methods("DELETE").Name("DeleteStudent")
}

func createRoutersTeacher(routerWS *mux.Router) {
	subRouter := routerWS.PathPrefix("/teacher").Subrouter()
	subRouter.Path("").HandlerFunc(service.FindAllTeacher).Methods("GET").Name("FindAllTeacher")
	subRouter.Path("/{id}").HandlerFunc(service.FindByIDTeacher).Methods("GET").Name("FindByIDTeacher")
	subRouter.Path("").HandlerFunc(service.InsertTeacher).Methods("POST").Name("InsertTeacher")
	subRouter.Path("").HandlerFunc(service.UpdateTeacher).Methods("PUT").Name("UpdateTeacher")
	subRouter.Path("/{id}").HandlerFunc(service.DeleteTeacher).Methods("DELETE").Name("DeleteTeacher")
}
