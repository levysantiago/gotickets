package users_services

import (
	users_models "github.com/levysantiago/gotickets/src/modules/users/models"
	users_repository "github.com/levysantiago/gotickets/src/modules/users/repositories"
)

func CreateUser(user *users_models.User) *users_models.User{
	return users_repository.Create(user)
}