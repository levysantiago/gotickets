package wallets_repository

import (
	wallets_models "github.com/levysantiago/gotickets/src/modules/wallets/models"
	db_provider "github.com/levysantiago/gotickets/src/shared/db"
)

var db = db_provider.GetDB()

func Create(wallet *wallets_models.Wallet) *wallets_models.Wallet{
	db.Create(wallet)
	return wallet
}

func Find(id string) *wallets_models.Wallet{
	wallet := wallets_models.Wallet{}
	db.Where("Id=?", id).First(&wallet)
	return &wallet
}

func FindByAddress(address string) *wallets_models.Wallet{
	wallet := wallets_models.Wallet{}
	db.Where("Address=?", address).First(&wallet)
	return &wallet
}