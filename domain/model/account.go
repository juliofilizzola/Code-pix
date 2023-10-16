package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Account struct {
	Base      `valid:"required"`
	OwnerName string    `json:"owner_name" valid:"notnull"`
	Bank      *Bank     `valid:"-"`
	Number    string    `json:"number" valid:"notnull"`
	PixKeys   []*PixKey `valid:"-"`
}

func (a Account) isValid() error {
	if _, err := govalidator.ValidateStruct(a); err != nil {
		return err
	}
	return nil
}

func NewAccount(bank *Bank, name string, number string) (*Account, error) {
	account := Account{
		OwnerName: name,
		Bank:      bank,
		Number:    number,
	}

	account.ID = uuid.NewV4().String()
	account.CreatedAt = time.Now()

	if err := account.isValid(); err != nil {
		return nil, err
	}

	return &account, nil
}
