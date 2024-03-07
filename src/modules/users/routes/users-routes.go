package users_routes

import (
	"github.com/gorilla/mux"
	users_controllers "github.com/levysantiago/gotickets/src/modules/users/controllers"
)

func Register(router *mux.Router){
	router.HandleFunc("/users/", users_controllers.CreateUserController).Methods("POST")
}