package route

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/felipesulzbach/example-go-api-rest/src/main/config"
	"github.com/gorilla/mux"

)

// LoadRouteServer Creates the HTTP server with the routes.
func LoadRouteServer() *http.Server {
	configs := config.GetConfig()
	routerWS := *mux.NewRouter()
	createRouters(&routerWS)

	domain, port := configs["api.domain"], configs["api.port"]
	appDomain := fmt.Sprintf("%s:%s", domain, port)

	log.Fatal(http.ListenAndServe(appDomain, &routerWS))

	srv := &http.Server{
		Addr:         appDomain,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      &routerWS,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	return srv
}
