package interfaces

import (
	"context"
	"e-wallet-wallet/internal/models"
	"github.com/gin-gonic/gin"
)

type IWalletRepository interface {
	CreateWallet(ctx context.Context, wallet *models.Wallet) error
	UpdateBalance(ctx context.Context, userID int, amount float64) (*models.Wallet, error)
	CreateWalletTrx(ctx context.Context, walletTrx *models.WalletTransaction) error
	GetWalletTransactionByReference(ctx context.Context, reference string) (*models.WalletTransaction, error)
	GetWalletByUserID(ctx context.Context, userID int) (*models.Wallet, error)
	GetWalletByID(ctx context.Context, walletID int) (*models.Wallet, error)
	GetWalletHistory(ctx context.Context, walletID int, offset int, limit int, transactionType string) ([]models.WalletTransaction, error)

	InsertWalletLink(ctx context.Context, req *models.WalletLink) error
	GetWalletLink(ctx context.Context, walletID int, clientSource string) (*models.WalletLink, error)
	UpdateStatusWalletLink(ctx context.Context, walletID int, clientSource string, status string) error
}

type IWalletService interface {
	Create(ctx context.Context, wallet *models.Wallet) error
	CreditBalance(ctx context.Context, userID int, req *models.TransactionRequest) (*models.BalanceResponse, error)
	DebitBalance(ctx context.Context, userID int, req *models.TransactionRequest) (*models.BalanceResponse, error)
	GetBalance(ctx context.Context, userID int) (*models.BalanceResponse, error)
	GetWalletHistory(ctx context.Context, userID int, param models.WalletHistoryParam) ([]models.WalletTransaction, error)

	ExGetBalance(ctx context.Context, walletID int) (*models.BalanceResponse, error)
	CreateWalletLink(ctx context.Context, clientSource string, req *models.WalletLink) (*models.WalletStructOTP, error)
	WalletLinkConfirmation(ctx context.Context, walletID int, clientSource string, otp string) error
	WalletUnlink(ctx context.Context, walletID int, clientSource string) error
}

type IWalletHandler interface {
	Create(c *gin.Context)
	CreditBalance(c *gin.Context)
	DebitBalance(c *gin.Context)
	GetBalance(c *gin.Context)
	GetWalletHistory(c *gin.Context)

	ExGetBalance(c *gin.Context)
	CrateWalletLink(c *gin.Context)
	WalletLinkConfirmation(c *gin.Context)
	WalletUnlink(c *gin.Context)
}
