package services

import (
	"context"
	"e-wallet-wallet/internal/interfaces"
	"e-wallet-wallet/internal/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"math/rand"
	"strconv"
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

func (s *WalletService) ExGetBalance(ctx context.Context, walletID int) (*models.BalanceResponse, error) {
	var (
		resp = &models.BalanceResponse{}
	)
	wallet, err := s.WalletRepository.GetWalletByUserID(ctx, walletID)
	if err != nil {
		return resp, errors.Wrap(err, "failed to get wallet by wallet id")
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

func (s *WalletService) CreateWalletLink(ctx context.Context, clientSource string, req *models.WalletLink) (*models.WalletStructOTP, error) {
	req.ClientSource = clientSource
	req.Status = "pending"
	req.OTP = strconv.Itoa(rand.Intn(999999))

	resp := &models.WalletStructOTP{
		OTP: req.OTP,
	}

	err := s.WalletRepository.InsertWalletLink(ctx, req)
	if err != nil {
		return resp, errors.Wrap(err, "failed to insert wallet link")
	}

	return resp, nil
}

func (s *WalletService) WalletLinkConfirmation(ctx context.Context, walletID int, clientSource string, otp string) error {
	walletLink, err := s.WalletRepository.GetWalletLink(ctx, walletID, clientSource)
	if err != nil {
		return errors.Wrap(err, "failed to get wallet link")
	}

	if walletLink.Status != "pending" {
		return errors.New("wallet link already confirmed")
	}

	if walletLink.OTP != otp {
		return errors.New("invalid otp")
	}

	return s.WalletRepository.UpdateStatusWalletLink(ctx, walletID, clientSource, "linked")
}

func (s *WalletService) WalletUnlink(ctx context.Context, walletID int, clientSource string) error {
	return s.WalletRepository.UpdateStatusWalletLink(ctx, walletID, clientSource, "unlinked")
}
