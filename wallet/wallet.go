package wallet

import (
	"fmt"

	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
)

type Wallet struct {
	ks *keystore.KeyStore
}

func NewWallet(keydir string) *Wallet {
	ks := keystore.NewKeyStore(keydir, keystore.LightScryptN, keystore.LightScryptP)
	return &Wallet{ks}
}

func (w *Wallet) NewAccount(pass string) (string, []byte, error) {
	account, err := w.ks.NewAccount(pass)
	keyJson, err := w.ks.Export(account, pass, pass)
	addr := account.Address.Hex()
	w.ks.Delete(account, pass)
	fmt.Println(string(keyJson))
	return addr, keyJson, err
}

func (w *Wallet) ImportAccount(keyJson []byte, pass string) (string, error) {
	account, err := w.ks.Import(keyJson, pass, pass)
	if err != nil {
		fmt.Println("failed to import account", err)
		return "", err
	}

	defer w.ks.Delete(account, pass)

	return account.Address.String(), nil

}

func (w Wallet) GetTransactOpts(keyJson []byte, pass string) (*bind.TransactOpts, error) {
	//addr := common.HexToAddress(address)

	account, err := w.ks.Import(keyJson, pass, pass)
	if err != nil {
		fmt.Println("failed to import account", err)
		return nil, err
	}
	defer w.ks.Delete(account, pass)
	w.ks.Unlock(account, pass)

	auth, err := bind.NewKeyStoreTransactor(w.ks, account)
	//defer w.ks.Lock(account.Address)

	return auth, err
}

func (w *Wallet) DeleteAccount(account accounts.Account, pass string) error {

	return w.ks.Delete(account, pass)
}

func (w Wallet) GetCallOpts(address string) *bind.CallOpts {
	addr := common.HexToAddress(address)
	return &bind.CallOpts{
		From: addr,
	}
}
