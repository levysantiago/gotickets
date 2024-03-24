package users_models

import (
	"time"

	"gorm.io/gorm"
)

type User struct{
	Id string `gorm:"type:uuid;default:gen_random_uuid()"`
	Email string `gorm:"unique"`
	Name string
	Password string `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index"json:"-"`
}