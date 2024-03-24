package users_controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	users_models "github.com/levysantiago/gotickets/src/modules/users/models"
	users_services "github.com/levysantiago/gotickets/src/modules/users/services"
)

type FindUserResponseData struct{
	Status string
	Data users_models.User
}

type FindUserErrorData struct{
	Status string
	ErrorMessage string
}

func FindUser(res http.ResponseWriter, req *http.Request){
	// Get id from params
	id := mux.Vars(req)["id"]

	// Find user
	user := users_services.FindUser(id)
	if(user.Id == ""){
		errorData := FindUserErrorData{Status: "error", ErrorMessage: "User not found."}
		data, _ := json.Marshal(&errorData)
		res.WriteHeader(http.StatusNotFound)
		res.Write(data)
		return
	}

	data, _ := json.Marshal(&user)
	res.WriteHeader(http.StatusNotFound)
	res.Write(data)
}

