package models

import "time"

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
	ID                    int       `gorm:"column:id;primary_key"`
	WalletID              int       `gorm:"column:wallet_id"`
	Amount                float64   `gorm:"column:amount;type:decimal(15,2)"`
	WalletTransactionType string    `gorm:"column:wallet_transaction_type;type:ENUM('CREDIT', 'DEBIT')"`
	Reference             string    `gorm:"column:reference;type:varchar(100);unique"`
	CreatedAt             time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdateAt              time.Time `gorm:"column:update_at;autoUpdateTime"`
}

func (*WalletTransaction) TableName() string {
	return "wallet_transactions"
}
