package interfaces

import (
	"context"
	"e-wallet-wallet/internal/models"
)

type IExternal interface {
	ValidateToken(ctx context.Context, token string) (*models.TokenData, error)
}
