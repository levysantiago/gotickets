package web3_provider

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"encoding/base64"
	"errors"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type Account struct{
	PubKey string
	PrivKey string
}

func CreateAccount() (Account, error){
	pk, err := crypto.GenerateKey()
	if(err != nil){
		return Account{}, err
	}

	pubKey := pk.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
  	err = errors.New("could not cast public key")
		return Account{}, err
  }
	
	privKeyBytes := crypto.FromECDSA(pk)
	

	account := Account{
		PubKey: crypto.PubkeyToAddress(*pubKeyECDSA).Hex(), 
		PrivKey: hexutil.Encode(privKeyBytes)[2:],
	}

	return account, nil
}

func (account *Account) Encrypt(secret string, iv string) (error){
	block, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return err
	}
	plainText := []byte(account.PrivKey)
	cfb := cipher.NewCFBEncrypter(block, []byte(iv))
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	account.PrivKey = base64.StdEncoding.EncodeToString(cipherText)
	return nil
}

func (account *Account) Decrypt(secret string, iv string) (error) {
	block, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return err
	}
	cipherText, err := base64.StdEncoding.DecodeString(account.PrivKey)
	if err != nil {
		return err
	}
	cfb := cipher.NewCFBDecrypter(block, []byte(iv))
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	account.PrivKey = string(plainText)
	return nil
}

func SignTx(tx *types.Transaction, pk string) (*types.Transaction, error){
	bn := new(big.Int)
	chainId, ok := bn.SetString(os.Getenv("CHAIN_ID"), 10)
	if(!ok){
		err := errors.New("invalid chainId on .env")
		return &types.Transaction{}, err
	}

	privateKey, errPk := crypto.HexToECDSA(pk)
	if(errPk != nil){
		return &types.Transaction{}, errPk
	}

	signedTx, errSign := types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)
	if(errSign != nil){
		return &types.Transaction{}, errSign
	}

	return signedTx, nil
}