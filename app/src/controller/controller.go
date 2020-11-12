package controller

import (
	"encoding/json"
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

func jsonOkResponse(w http.ResponseWriter, value interface{}) {
	jsonResponse(w, value, 200)
}

func jsonCreatedResponse(w http.ResponseWriter, value interface{}) {
	jsonResponse(w, value, 201)
}

func jsonResponse(w http.ResponseWriter, value interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	if value != "" {
		json.NewEncoder(w).Encode(value)
	}
}
