package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/felipesulzbach/example-go-api-rest/src/infra/db"
	"github.com/felipesulzbach/example-go-api-rest/src/main/config"
	"github.com/felipesulzbach/example-go-api-rest/src/main/route"
	"github.com/felipesulzbach/example-go-api-rest/src/migrate"

)

func main() {
	log.Println("SERVER Starting...")

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "he duration for which the server normally expects existing connections to end - e.g. 15s or 1m")
	flag.Parse()

	db.ConnectionTest()
	config.LoadFileConfig()
	migrate.LoadDatabaseStructure()

	log.Println("SERVER Successfully started!")

	routeServer := route.LoadRouteServer()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	routeServer.Shutdown(ctx)
	log.Println("SERVER Shutting Down!")
	os.Exit(0)
}
