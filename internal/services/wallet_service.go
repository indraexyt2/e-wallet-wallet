package services

import (
	"context"
	"e-wallet-wallet/internal/interfaces"
	"e-wallet-wallet/internal/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type WalletService struct {
	WalletRepository interfaces.IWalletRepository
}

func (s *WalletService) Create(ctx context.Context, wallet *models.Wallet) error {
	return s.WalletRepository.CreateWallet(ctx, wallet)
}

func (s *WalletService) CreditBalance(ctx context.Context, userID int, req *models.TransactionRequest) (*models.BalanceResponse, error) {
	var (
		resp models.BalanceResponse
	)

	history, err := s.WalletRepository.GetWalletTransactionByReference(ctx, req.Reference)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return &resp, errors.Wrap(err, "failed to get wallet transaction by reference")
		}
	}

	if history.ID > 0 {
		return &resp, errors.New("wallet transaction already exists")
	}

	wallet, err := s.WalletRepository.UpdateBalance(ctx, userID, req.Amount)
	if err != nil {
		return &resp, errors.Wrap(err, "failed to update balance")
	}

	walletTrx := &models.WalletTransaction{
		WalletID:              wallet.ID,
		Amount:                req.Amount,
		WalletTransactionType: "CREDIT",
		Reference:             req.Reference,
	}

	err = s.WalletRepository.CreateWalletTrx(ctx, walletTrx)
	if err != nil {
		return &resp, errors.Wrap(err, "failed to create wallet transaction")
	}

	resp.Balance = wallet.Balance + req.Amount

	return &resp, nil
}

func (s *WalletService) DebitBalance(ctx context.Context, userID int, req *models.TransactionRequest) (*models.BalanceResponse, error) {
	var (
		resp models.BalanceResponse
	)

	history, err := s.WalletRepository.GetWalletTransactionByReference(ctx, req.Reference)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return &resp, errors.Wrap(err, "failed to get wallet transaction by reference")
		}
	}

	if history.ID > 0 {
		return &resp, errors.New("wallet transaction already exists")
	}

	wallet, err := s.WalletRepository.UpdateBalance(ctx, userID, -req.Amount)
	if err != nil {
		return &resp, errors.Wrap(err, "failed to update balance")
	}

	walletTrx := &models.WalletTransaction{
		WalletID:              wallet.ID,
		Amount:                req.Amount,
		WalletTransactionType: "DEBIT",
		Reference:             req.Reference,
	}

	err = s.WalletRepository.CreateWalletTrx(ctx, walletTrx)
	if err != nil {
		return &resp, errors.Wrap(err, "failed to create wallet transaction")
	}

	resp.Balance = wallet.Balance - req.Amount

	return &resp, nil
}

func (s *WalletService) GetBalance(ctx context.Context, userID int) (*models.BalanceResponse, error) {
	var (
		resp = &models.BalanceResponse{}
	)
	wallet, err := s.WalletRepository.GetWalletByUserID(ctx, userID)
	if err != nil {
		return resp, errors.Wrap(err, "failed to get wallet by user id")
	}

	resp.Balance = wallet.Balance

	return resp, nil
}

func (s *WalletService) GetWalletHistory(ctx context.Context, userID int, param models.WalletHistoryParam) ([]models.WalletTransaction, error) {
	var (
		resp []models.WalletTransaction
	)

	wallet, err := s.WalletRepository.GetWalletByUserID(ctx, userID)
	if err != nil {
		return resp, errors.Wrap(err, "failed to get wallet by user id")
	}

	offset := (param.Page - 1) * param.Limit
	resp, err = s.WalletRepository.GetWalletHistory(ctx, wallet.ID, offset, param.Limit, param.WalletTransactionType)
	if err != nil {
		return resp, errors.Wrap(err, "failed to get wallet history")
	}

	return resp, nil
}
