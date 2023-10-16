package model

type Transaction struct {
	Base              `valid:"required"`
	AccountFrom       *Account `json:"account_from" valid:"-"`
	Amount            float64  `json:"amount" valid:"notnull"`
	PixKeyTo          *PixKey  `json:"pix_key_to" valid:"-"`
	Status            string   `json:"status" valid:"notnull"`
	Description       string   `json:"description" valid:"notnull"`
	CancelDescription string   `json:"cancel_description" valid:"notnull"`
}
