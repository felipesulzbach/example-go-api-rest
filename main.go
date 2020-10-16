package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/felipesulzbach/exemplo-api-rest/app/src/controller"
	"github.com/felipesulzbach/exemplo-api-rest/app/src/migrate"
	"github.com/felipesulzbach/exemplo-api-rest/app/src/repository"
	"github.com/gorilla/mux"

)

// https://medium.com/@ishagirdhar/a-simple-restful-web-server-using-golang-mysql-gorilla-mux-53f1568d22e5

func main() {
	log.Println("SERVER Starting...")
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "he duration for which the server normally expects existing connections to end - e.g. 15s or 1m")
	flag.Parse()

	repository.TestConnectionDB()
	migrate.LoadDatabaseStructure()
	srv := bootHTTPServer()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("SERVER Shutting Down!")
	os.Exit(0)
}

// bootHTTPServer - Creates the HTTP server with the routes.
func bootHTTPServer() *http.Server {
	routerWS := *mux.NewRouter()
	controller.CreateRouters(&routerWS)

	log.Println("SERVER Successfully started!")
	log.Fatal(http.ListenAndServe(controller.Domain, &routerWS))

	srv := &http.Server{
		Addr:         controller.Domain,
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
