package models

import "time"

type Wallet struct {
	ID        int       `gorm:"column:id;primary_key"`
	UserID    int       `gorm:"column:user_id"`
	Balance   int       `gorm:"column:balance;type:decimal(15,2)"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdateAt  time.Time `gorm:"column:update_at;autoUpdateTime"`
}

func (*Wallet) TableName() string {
	return "wallets"
}

type WalletTransaction struct {
	ID                    int       `gorm:"column:id;primary_key"`
	WalletID              int       `gorm:"column:wallet_id"`
	Amount                int       `gorm:"column:amount;type:decimal(15,2)"`
	WalletTransactionType string    `gorm:"column:amount;type:ENUM('CREDIT', 'DEBIT')"`
	Reference             string    `gorm:"column:reference;type:varchar(100)"`
	CreatedAt             time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdateAt              time.Time `gorm:"column:update_at;autoUpdateTime"`
}

func (*WalletTransaction) TableName() string {
	return "wallet_transactions"
}
