package users_repository

import (
	users_models "github.com/levysantiago/gotickets/src/modules/users/models"
	db_provider "github.com/levysantiago/gotickets/src/shared/db"
)

var db = db_provider.GetDB()

func Create(user *users_models.User) *users_models.User{
	db.Create(user)
	return user
}

func Find(id string) *users_models.User{
	user := users_models.User{}
	db.Where("Id=?", id).First(&user)
	return &user
}

func FindByEmail(email string) *users_models.User{
	user := users_models.User{}
	db.Where("Email=?", email).First(&user)
	return &user
}

func FindMany() []users_models.User{
	var users []users_models.User
	db.Find(&users)
	return users
}

func Update(user *users_models.User) *users_models.User{
	db.Where("Id=?", user.Id).Updates(&user)
	return user
}