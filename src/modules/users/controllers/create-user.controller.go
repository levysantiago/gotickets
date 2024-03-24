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
	// User data
	var createUserData = users_models.User{}

	// Decoding body
	var err = json.NewDecoder(req.Body).Decode(&createUserData)
	if(err != nil){
		var errorData = CreateUserErrorData{Status: "error", ErrorMessage: "Failed to decode body."}
		var data, _ = json.Marshal(errorData)
		res.WriteHeader(http.StatusInternalServerError)
		res.Write(data)
		return
	}

	// Creating user
	_, err = users_services.CreateUser(&createUserData)
	if(err != nil){
		errorData := CreateUserErrorData{Status: "error", ErrorMessage: err.Error()}
		data, _ := json.Marshal(errorData)
		res.WriteHeader(http.StatusBadRequest)
		res.Write(data)
		return
	}

	var data, _ = json.Marshal(createUserData)
	res.WriteHeader(http.StatusCreated)
	res.Write(data)
}