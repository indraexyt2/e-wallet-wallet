package repository

import (
	"context"
	"e-wallet-wallet/internal/models"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type WalletRepository struct {
	DB *gorm.DB
}

func (r *WalletRepository) CreateWallet(ctx context.Context, wallet *models.Wallet) error {
	return r.DB.Create(wallet).Error
}

func (r *WalletRepository) UpdateBalance(ctx context.Context, userID int, amount float64) (*models.Wallet, error) {
	var wallet *models.Wallet
	err := r.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&models.Wallet{}).
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Select("balance").
			Where("user_id =?", userID).
			Take(&wallet).Error

		if (wallet.Balance + amount) < 0 {
			return fmt.Errorf("insufficient balance")
		}

		err = tx.Model(&models.Wallet{}).
			Where("user_id = ?", userID).
			Update("balance", gorm.Expr("balance + ?", amount)).Error

		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to update balance: %w", err)
	}

	return wallet, nil
}

func (r *WalletRepository) CreateWalletTrx(ctx context.Context, walletTrx *models.WalletTransaction) error {
	return r.DB.Create(walletTrx).Error
}

func (r *WalletRepository) GetWalletTransactionByReference(ctx context.Context, reference string) (*models.WalletTransaction, error) {
	var (
		resp models.WalletTransaction
	)
	err := r.DB.Where("reference = ?", reference).Last(&resp).Error
	if err != nil {
		return &resp, err
	}

	return &resp, nil
}

func (r *WalletRepository) GetWalletByUserID(ctx context.Context, userID int) (*models.Wallet, error) {
	var (
		resp models.Wallet
	)
	err := r.DB.Where("user_id = ?", userID).Last(&resp).Error
	if err != nil {
		return &resp, err
	}

	return &resp, nil
}

func (r *WalletRepository) GetWalletByID(ctx context.Context, walletID int) (*models.Wallet, error) {
	var (
		resp models.Wallet
	)
	err := r.DB.Where("user_id = ?", walletID).Last(&resp).Error
	if err != nil {
		return &resp, err
	}

	return &resp, nil
}

func (r *WalletRepository) GetWalletHistory(ctx context.Context, walletID int, offset int, limit int, transactionType string) ([]models.WalletTransaction, error) {
	var (
		resp []models.WalletTransaction
	)

	sql := r.DB
	if transactionType != "" {
		sql = sql.Where("wallet_transaction_type = ?", transactionType)
	}

	err := sql.Model(&models.WalletTransaction{}).Limit(limit).Offset(offset).Order("id DESC").Find(&resp).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get wallet history: %w", err)
	}

	fmt.Println(resp)

	return resp, nil
}

func (r *WalletRepository) InsertWalletLink(ctx context.Context, req *models.WalletLink) error {
	return r.DB.Create(req).Error
}

func (r *WalletRepository) GetWalletLink(ctx context.Context, walletID int, clientSource string) (*models.WalletLink, error) {
	var (
		resp models.WalletLink
		err  error
	)
	err = r.DB.Where("wallet_id = ?", walletID).Where("client_source = ?", clientSource).First(&resp).Error
	if err != nil {
		return &resp, err
	}
	return &resp, nil
}

func (r *WalletRepository) UpdateStatusWalletLink(ctx context.Context, walletID int, clientSource string, status string) error {
	return r.DB.Model(&models.WalletLink{}).Where("wallet_id = ?", walletID).Where("client_source = ?", clientSource).Update("status", status).Error
}
