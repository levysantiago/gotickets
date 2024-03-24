package users_models

import (
	"time"

	wallets_models "github.com/levysantiago/gotickets/src/modules/wallets/models"
)

type User struct{
	Id string `gorm:"type:uuid;default:gen_random_uuid()"`
	Email string `gorm:"unique"`
	Name string
	WalletId string
	Wallet wallets_models.Wallet `gorm:"foreignkey:WalletId;references:Id;`
	Password string 
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}