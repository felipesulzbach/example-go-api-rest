package route

import (
	"log"
	"net/http"

	"github.com/felipesulzbach/example-go-api-rest/src/main/config"
	"github.com/felipesulzbach/example-go-api-rest/src/presentation/controller"
	"github.com/gorilla/mux"

)

// Create routes of Endpoints
func createRouters(routerWS *mux.Router) {
	log.Println("ROUTERS Creating...")

	_createRoutersCourse(routerWS)
	//_createRoutersSchoolClass(routerWS)
	//_createRoutersStudent(routerWS)
	//_createRoutersTeacher(routerWS)
	http.Handle("/", routerWS)

	log.Println("ROUTERS Successfully created!")
}

func printRoutes(router *mux.Router, name string) {
	configs := config.GetConfig()

	log.Printf("%s:\n", name)
	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		template, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		log.Printf("> %s://%s:%s%s", configs["api.protocol"], configs["api.domain"], configs["api.port"], template)
		return nil
	})
}

func _createRoutersCourse(routerWS *mux.Router) {
	subRouter := routerWS.PathPrefix("/api/v1/course").Subrouter()
	subRouter.Path("").HandlerFunc(controller.GetAllCourse).Methods("GET").Name("FindAllCourse")
	subRouter.Path("/{id}").HandlerFunc(controller.GetByIDCourse).Methods("GET").Name("FindByIDCourse")
	subRouter.Path("").HandlerFunc(controller.CreateCourse).Methods("POST").Name("CreateCourse")
	subRouter.Path("").HandlerFunc(controller.UpdateCourse).Methods("PUT").Name("UpdateCourse")
	subRouter.Path("/{id}").HandlerFunc(controller.DeleteCourse).Methods("DELETE").Name("DeleteCourse")

	printRoutes(subRouter, "Course")
}

/*func _createRoutersSchoolClass(routerWS *mux.Router) {
	subRouter := routerWS.PathPrefix("/api/v1/schoolclass").Subrouter()
	subRouter.Path("").HandlerFunc(controller.GetAllSchoolClass).Methods("GET").Name("FindAllSchoolClass")
	subRouter.Path("/{id}").HandlerFunc(controller.GetByIDSchoolClass).Methods("GET").Name("FindByIDSchoolClass")
	subRouter.Path("").HandlerFunc(controller.CreateSchoolClass).Methods("POST").Name("CreateSchoolClass")
	subRouter.Path("").HandlerFunc(controller.UpdateSchoolClass).Methods("PUT").Name("UpdateSchoolClass")
	subRouter.Path("/{id}").HandlerFunc(controller.DeleteSchoolClass).Methods("DELETE").Name("DeleteSchoolClass")

	printRoutes(subRouter, "SchoolClass")
}

func _createRoutersStudent(routerWS *mux.Router) {
	subRouter := routerWS.PathPrefix("/api/v1/student").Subrouter()
	subRouter.Path("").HandlerFunc(controller.GetAllStudent).Methods("GET").Name("FindAllStudent")
	subRouter.Path("/{id}").HandlerFunc(controller.GetByIDStudent).Methods("GET").Name("FindByIDStudent")
	subRouter.Path("").HandlerFunc(controller.CreateStudent).Methods("POST").Name("CreateStudent")
	subRouter.Path("").HandlerFunc(controller.UpdateStudent).Methods("PUT").Name("UpdateStudent")
	subRouter.Path("/{id}").HandlerFunc(controller.DeleteStudent).Methods("DELETE").Name("DeleteStudent")

	printRoutes(subRouter, "Student")
}

func _createRoutersTeacher(routerWS *mux.Router) {
	subRouter := routerWS.PathPrefix("/api/v1/teacher").Subrouter()
	subRouter.Path("").HandlerFunc(controller.GetAllTeacher).Methods("GET").Name("FindAllTeacher")
	subRouter.Path("/{id}").HandlerFunc(controller.GetByIDTeacher).Methods("GET").Name("FindByIDTeacher")
	subRouter.Path("").HandlerFunc(controller.CreateTeacher).Methods("POST").Name("CreateTeacher")
	subRouter.Path("").HandlerFunc(controller.UpdateTeacher).Methods("PUT").Name("UpdateTeacher")
	subRouter.Path("/{id}").HandlerFunc(controller.DeleteTeacher).Methods("DELETE").Name("DeleteTeacher")

	printRoutes(subRouter, "Teacher")
}*/
