package models

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type Wallet struct {
	ID        int       `gorm:"column:id;primary_key" json:"id"`
	UserID    int       `gorm:"column:user_id;unique" json:"user_id"`
	Balance   float64   `gorm:"column:balance;type:decimal(15,2)" json:"balance"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"-"`
	UpdateAt  time.Time `gorm:"column:update_at;autoUpdateTime" json:"-"`
}

func (*Wallet) TableName() string {
	return "wallets"
}

type WalletTransaction struct {
	ID                    int       `json:"-" gorm:"column:id;primary_key"`
	WalletID              int       `json:"wallet_id" gorm:"column:wallet_id"`
	Amount                float64   `json:"amount" gorm:"column:amount;type:decimal(15,2)"`
	WalletTransactionType string    `json:"wallet_transaction_type" gorm:"column:wallet_transaction_type;type:ENUM('CREDIT', 'DEBIT')"`
	Reference             string    `json:"reference" gorm:"column:reference;type:varchar(100);unique"`
	CreatedAt             time.Time `json:"date" gorm:"column:created_at;autoCreateTime"`
	UpdateAt              time.Time `json:"-" gorm:"column:update_at;autoUpdateTime"`
}

func (*WalletTransaction) TableName() string {
	return "wallet_transactions"
}

type WalletHistoryParam struct {
	Page                  int    `form:"page"`
	Limit                 int    `form:"limit"`
	WalletTransactionType string `form:"wallet_transaction_type"`
}

type WalletLink struct {
	ID           int       `json:"id"`
	WalletID     int       `gorm:"column:wallet_id" json:"wallet_id" validate:"required"`
	ClientSource string    `gorm:"column:client_source;type:varchar(100)" json:"client_source"`
	OTP          string    `gorm:"column:otp;type:varchar(6)" json:"otp"`
	Status       string    `gorm:"column:status;type:varchar(10)" json:"status"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
}

func (*WalletLink) TableName() string {
	return "wallet_links"
}

func (w *WalletLink) Validate() error {
	v := validator.New()
	return v.Struct(w)
}

type WalletStructOTP struct {
	OTP string `json:"otp" validate:"required"`
}
