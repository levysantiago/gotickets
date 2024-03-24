package users_controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	users_services "github.com/levysantiago/gotickets/src/modules/users/services"
)

type UpdateUserErrorData struct{
	Status string
	ErrorMessage string
}

func UpdateUser(res http.ResponseWriter, req *http.Request){
	userId := mux.Vars(req)["id"]
	
	var dataToUpdate users_services.UpdateUserData
	json.NewDecoder(req.Body).Decode(&dataToUpdate)

	userUpdated, err := users_services.UpdateUser(userId, &dataToUpdate)
	if(err != nil){
		errorData := UpdateUserErrorData{Status: "error", ErrorMessage: err.Error()} 
		data, _ := json.Marshal(&errorData)
		res.WriteHeader(http.StatusBadRequest)
		res.Write(data)
		return
	}

	data, _ := json.Marshal(&userUpdated)
	res.WriteHeader(http.StatusBadRequest)
	res.Write(data)
}