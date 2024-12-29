package interfaces

import (
	"context"
	"e-wallet-wallet/internal/models"
	"github.com/gin-gonic/gin"
)

type IWalletRepository interface {
	CreateWallet(ctx context.Context, wallet *models.Wallet) error
}

type IWalletService interface {
	Create(ctx context.Context, wallet *models.Wallet) error
}

type IWalletHandler interface {
	Create(c *gin.Context)
}
