package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Bank struct {
	Id        string    `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewBank(code string, name string) (*Bank, error) {
	bank := Bank{
		Code: code,
		Name: name,
	}

	bank.Id = uuid.NewV4().String()
	bank.CreatedAt = time.Now()

	return &bank, nil
}
