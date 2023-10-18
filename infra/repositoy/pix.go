package repositoy

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/juliofilizzola/Code-pix/domain/model"
)

type PixRepositoryDb struct {
	Db *gorm.DB
}

func (receiver PixRepositoryDb) AddBank(bank *model.Bank) error {
	if err := receiver.Db.Create(bank).Error; err != nil {
		return err
	}
	return nil
}

func (receiver PixRepositoryDb) AddAccount(account *model.Account) error {
	if err := receiver.Db.Create(account).Error; err != nil {
		return err
	}
	return nil
}

func (receiver PixRepositoryDb) RegisterKey(pixKey *model.PixKey) (*model.PixKey, error) {
	if err := receiver.Db.Create(pixKey).Error; err != nil {
		return nil, err
	}
	return pixKey, nil
}

func (receiver PixRepositoryDb) FindKeyByKind(key string, kind string) (*model.PixKey, error) {
	var pixKey model.PixKey
	receiver.Db.Preload("Account.Bank").First(&pixKey, "kind = ? and key = ?", kind, key)

	if pixKey.ID == "" {
		return nil, fmt.Errorf("no key was found")
	}
	return &pixKey, nil
}

func (receiver PixRepositoryDb) FindAccount(id string) (*model.Account, error) {
	var account model.Account
	receiver.Db.Preload("Bank").First(&account, "id = ?", id)

	if account.ID == "" {
		return nil, fmt.Errorf("no account found")
	}
	return &account, nil
}

func (receiver PixRepositoryDb) FindBank(id string) (*model.Bank, error) {
	var bank model.Bank
	receiver.Db.First(&bank, "id = ?", id)

	if bank.ID == "" {
		return nil, fmt.Errorf("no bank found")
	}
	return &bank, nil
}
