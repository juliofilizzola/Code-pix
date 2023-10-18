package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type PixKeyRepositoryInterface interface {
	Register(pixKey *PixKey) (*PixKey, error)
	FindKeyByKind(key string, kind string) (*PixKey, error)
	AddBank(bank *Bank) error
	AddAccount(account *Account) error
	FindAccount(id string) (*Account, error)
}

type Status int

// Declare related constants for each weekday starting with index 1
const (
	Active   Status = iota + 1 // EnumIndex = 1
	Inactive                   // EnumIndex = 2 	// EnumIndex = 7
)

type PixKey struct {
	Base      `valid:"required"`
	Kind      string   `json:"kind" gorm:"type:varchar(20)" valid:"notnull"`
	Key       string   `json:"key" gorm:"type:varchar(255)" valid:"notnull"`
	AccountID string   `json:"account_id" gorm:"column:account_id;type:uuid;not null" valid:"notnull"`
	Account   *Account `valid:"-"`
	Status    Status   `json:"status" gorm:"type:varchar(20)" valid:"notnull"`
}

func (p PixKey) isValid() error {
	if _, err := govalidator.ValidateStruct(p); err != nil {
		return err
	}
	if p.Kind != "email" && p.Kind != "cpf" {
		return errors.New("invalid type key")
	}
	if p.Status != Active && p.Status != Inactive {
		return errors.New("invalid status")

	}
	return nil
}

func NewPikKey(kind string, key string, account *Account) (*PixKey, error) {
	pixKey := PixKey{
		Account: account,
		Kind:    kind,
		Key:     key,
		Status:  Active,
	}

	pixKey.ID = uuid.NewV4().String()
	pixKey.CreatedAt = time.Now()

	if err := pixKey.isValid(); err != nil {
		return nil, err
	}

	return &pixKey, nil
}

func (s Status) String() string {
	return [...]string{"active", "inactive"}[s-1]
}
