package services

import (
	"context"
	"e-wallet-wallet/internal/interfaces"
	"e-wallet-wallet/internal/models"
)

type WalletService struct {
	WalletRepository interfaces.IWalletRepository
}

func (s *WalletService) Create(ctx context.Context, wallet *models.Wallet) error {
	return s.WalletRepository.CreateWallet(ctx, wallet)
}
