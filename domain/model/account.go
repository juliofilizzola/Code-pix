package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Account struct {
	Base      `valid:"required"`
	OwnerName string    `json:"owner_name" gorm:"column:owner_name;type:varchar(255);not null" valid:"notnull"`
	Bank      *Bank     `valid:"-"`
	BankID    string    `json:"bank_id" gorm:"column:bank_id;type:uuid;type:not null" valid:"-"`
	Number    string    `json:"number" gorm:"type:varchar(20)" valid:"notnull"`
	PixKeys   []*PixKey `gorm:"ForeignKey:AccountID" valid:"-"`
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
