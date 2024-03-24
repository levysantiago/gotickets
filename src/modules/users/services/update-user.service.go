package users_services

import (
	"errors"

	users_models "github.com/levysantiago/gotickets/src/modules/users/models"
	users_repository "github.com/levysantiago/gotickets/src/modules/users/repositories"
)

type UpdateUserData struct{
	Name string
}

func UpdateUser(id string, dataToUpdate *UpdateUserData) (*users_models.User, error) {
	user := users_repository.Find(id)
	if(user.Id == ""){
		err := errors.New("user not found")
		var emptyUser users_models.User
		return &emptyUser, err
	}

	user.Name = dataToUpdate.Name

	users_repository.Update(user)
	return user, nil
}