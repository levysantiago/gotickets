package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	heartbeat_routes "github.com/levysantiago/gotickets/src/modules/heartbeat/routes"
)

func main(){
	var err = godotenv.Load(".env")

	if(err != nil){
		log.Fatal("Error loading .env file")
	}

	router := mux.NewRouter()

	heartbeat_routes.Register(router)

	http.Handle("/", router)

	fmt.Println("🚀 Server started at localhost:3333")
	log.Fatal(http.ListenAndServe("localhost:3333", router))
}