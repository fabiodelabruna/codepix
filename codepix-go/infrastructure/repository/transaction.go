package repository

import (
	"fmt"

	"github.com/fabiodelabruna/codepix/codepix-go/domain/model"
	"gorm.io/gorm"
)

type TransactionRepositoryDB struct {
	Db *gorm.DB
}

func (repository PixKeyRepositoryDB) Register(transaction *model.Transaction) error {
	err := repository.Db.Create(transaction).Error
	if err != nil {
		return err
	}
	return nil
}

func (repository PixKeyRepositoryDB) Save(transaction *model.Transaction) error {
	err := repository.Db.Save(transaction).Error
	if err != nil {
		return err
	}
	return nil
}

func (repository PixKeyRepositoryDB) Find(id string) (*model.Transaction, error) {
	var transaction model.Transaction
	repository.Db.Preload("AccountFrom.Bank").First(&transaction, "id = ?", id)

	if transaction.ID == "" {
		return nil, fmt.Errorf("no transaction was found")
	}

	return &transaction, nil
}
