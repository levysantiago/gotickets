package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	heartbeat_routes "github.com/levysantiago/gotickets/src/modules/heartbeat/routes"
	users_routes "github.com/levysantiago/gotickets/src/modules/users/routes"
	shared_middlewares "github.com/levysantiago/gotickets/src/shared/http/middlewares"
)

func main(){
	router := mux.NewRouter()

	// middlewares
	router.Use(shared_middlewares.ContentTypeMiddleware)

	// routes
	heartbeat_routes.Register(router)
	users_routes.Register(router)

	http.Handle("/", router)

	fmt.Println("🚀 Server started at localhost:3333")
	log.Fatal(http.ListenAndServe("localhost:3333", router))
}