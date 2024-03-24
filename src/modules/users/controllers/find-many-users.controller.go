package users_controllers

import (
	"encoding/json"
	"net/http"

	users_models "github.com/levysantiago/gotickets/src/modules/users/models"
	users_repository "github.com/levysantiago/gotickets/src/modules/users/repositories"
)

type FindManyUsersResponseData struct{
	Status string
	Data []users_models.User
}

func FindManyUsers(res http.ResponseWriter, req *http.Request){
	users := users_repository.FindMany()

	data, _ := json.Marshal(users)

	res.WriteHeader(http.StatusOK)
	res.Write(data)
}