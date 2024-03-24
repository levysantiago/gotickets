package users_services

import (
	"errors"

	users_models "github.com/levysantiago/gotickets/src/modules/users/models"
	users_repository "github.com/levysantiago/gotickets/src/modules/users/repositories"
)

func CreateUser(user *users_models.User) (*users_models.User, error){
	userWithSameEmail := users_repository.FindByEmail(user.Email)
	
	if(userWithSameEmail.Email != ""){	
		err := errors.New("email already registered")
		var emptyUser users_models.User
		return &emptyUser, err
	}
	
	userCreated := users_repository.Create(user)
	return userCreated, nil
}