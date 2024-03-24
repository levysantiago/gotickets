package users_controllers

import (
	"encoding/json"
	"net/http"

	users_models "github.com/levysantiago/gotickets/src/modules/users/models"
	users_services "github.com/levysantiago/gotickets/src/modules/users/services"
)

type CreateUserResponseData struct{
	Status string 
	Data users_models.User
}

type CreateUserErrorData struct{
	Status string 
	ErrorMessage string
}

func CreateUser(res http.ResponseWriter, req *http.Request){
	// Defining content type
	res.Header().Set("Content-Type", "application/json")

	// User data
	var createUserData = users_models.User{}

	// Decoding body
	var err = json.NewDecoder(req.Body).Decode(&createUserData)
	if(err != nil){
		var errorData = CreateUserErrorData{Status: "error", ErrorMessage: "Failed to decode body."}
		var data, _ = json.Marshal(errorData)
		res.WriteHeader(http.StatusInternalServerError)
		res.Write(data)
	}

	// Creating user
	var user = users_services.CreateUser(&createUserData)

	var data, _ = json.Marshal(user)
	res.WriteHeader(http.StatusCreated)
	res.Write(data)
}