package repository

import (
	"fmt"

	"github.com/fabiodelabruna/codepix/codepix-go/domain/model"
	"gorm.io/gorm"
)

type PixKeyRepositoryDB struct {
	Db *gorm.DB
}

func (repository PixKeyRepositoryDB) AddBank(bank *model.Bank) error {
	err := repository.Db.Create(bank).Error
	if err != nil {
		return err
	}
	return nil
}

func (repository PixKeyRepositoryDB) AddAccount(account *model.Account) error {
	err := repository.Db.Create(account).Error
	if err != nil {
		return err
	}
	return nil
}

func (repository PixKeyRepositoryDB) RegisterKey(pixKey *model.PixKey) error {
	err := repository.Db.Create(pixKey).Error
	if err != nil {
		return err
	}
	return nil
}

func (repository PixKeyRepositoryDB) FindKeyByKind(key string, kind string) (*model.PixKey, error) {
	var pixKey model.PixKey
	repository.Db.Preload("Account.Bank").First(&pixKey, "kind = ? and key = ?", kind, key)

	if pixKey.ID == "" {
		return nil, fmt.Errorf("no key was found")
	}

	return &pixKey, nil
}

func (repository PixKeyRepositoryDB) FindAccount(id string) (*model.Account, error) {
	var account model.Account
	repository.Db.Preload("Bank").First(&account, "id = ?", id)

	if account.ID == "" {
		return nil, fmt.Errorf("no account was found")
	}

	return &account, nil
}

func (repository PixKeyRepositoryDB) FindBank(id string) (*model.Bank, error) {
	var bank model.Bank
	repository.Db.First(&bank, "id = ?", id)

	if bank.ID == "" {
		return nil, fmt.Errorf("no bank was found")
	}

	return &bank, nil
}
