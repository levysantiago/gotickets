package users_models

import (
	"time"

	"gorm.io/gorm"
)

type User struct{
	Id string `gorm:"type:uuid;default:gen_random_uuid()"`
	Email string `gorm:"unique"`
	Password string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}