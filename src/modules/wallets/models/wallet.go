package wallets_models

import "time"

type Wallet struct{
	Id string `gorm:"type:uuid;default:gen_random_uuid()"`
	Address string `gorm:"unique"`
	PrivKey string 
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

