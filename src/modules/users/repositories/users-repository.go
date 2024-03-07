package users_repository

import (
	users_models "github.com/levysantiago/gotickets/src/modules/users/models"
	db_provider "github.com/levysantiago/gotickets/src/shared/db"
)

var db = db_provider.GetDB()

func CreateUser(user *users_models.User) *users_models.User{
	db.Create(user)
	return user
}