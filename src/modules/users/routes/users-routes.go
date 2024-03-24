package users_routes

import (
	"github.com/gorilla/mux"
	users_controllers "github.com/levysantiago/gotickets/src/modules/users/controllers"
)

func Register(router *mux.Router){
	router.HandleFunc("/users/", users_controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", users_controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", users_controllers.FindUser).Methods("GET")
	router.HandleFunc("/users", users_controllers.FindManyUsers).Methods("GET")
}