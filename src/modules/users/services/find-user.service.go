package users_services

import (
	users_models "github.com/levysantiago/gotickets/src/modules/users/models"
	users_repository "github.com/levysantiago/gotickets/src/modules/users/repositories"
)

func FindUser(id string) *users_models.User{
	return users_repository.Find(id)
}