package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

)

const (
	host = "localhost"
	port = "8080"

	// APIDomain ...
	APIDomain = host + ":" + port
)

// CreateRouters - Create routes to endpoints.
func CreateRouters(routerWS *mux.Router) {
	log.Println("ROUTERS Creating...")

	createRoutersCourse(routerWS)
	createRoutersSchoolClass(routerWS)
	createRoutersPerson(routerWS)
	createRoutersStudent(routerWS)
	createRoutersTeacher(routerWS)
	http.Handle("/", routerWS)

	log.Println("ROUTERS Successfully created!")
}

func printRoutes(router *mux.Router, name string) {
	log.Printf("%s:\n", name)
	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		template, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		log.Printf("> http://%s%s", APIDomain, template)
		return nil
	})
}
