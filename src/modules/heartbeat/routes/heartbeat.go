package heartbeat_routes

import (
	"github.com/gorilla/mux"
	heartbeat_controller "github.com/levysantiago/gotickets/src/modules/heartbeat/controllers"
)

func Register(router *mux.Router){
	router.HandleFunc("/", heartbeat_controller.Handle).Methods("GET")
}