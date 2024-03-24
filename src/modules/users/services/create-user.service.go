package users_services

import (
	"errors"
	"os"

	"github.com/google/uuid"
	users_models "github.com/levysantiago/gotickets/src/modules/users/models"
	users_repository "github.com/levysantiago/gotickets/src/modules/users/repositories"
	wallets_models "github.com/levysantiago/gotickets/src/modules/wallets/models"
	web3_provider "github.com/levysantiago/gotickets/src/shared/providers"
)

func CreateUser(user *users_models.User) (*users_models.User, error){
	userWithSameEmail := users_repository.FindByEmail(user.Email)
	
	if(userWithSameEmail.Email != ""){	
		err := errors.New("email already registered")
		return &users_models.User{}, err
	}

	account, err := web3_provider.CreateAccount()
	if(err != nil){
		return &users_models.User{}, err
	}

	secret := os.Getenv("WALLET_SECRET")
	iv := os.Getenv("WALLET_IV_BYTES")
	errEncrypt := account.Encrypt(secret, iv)
	if(errEncrypt != nil){
		return &users_models.User{}, errEncrypt
	}

	wallet := wallets_models.Wallet{
		Id: uuid.NewString(),
		Address: account.PubKey,
		PrivKey: account.PrivKey,
	}

	user.Id = uuid.NewString()
	user.Wallet = wallet
	
	userCreated := users_repository.Create(user)
	return userCreated, nil
}