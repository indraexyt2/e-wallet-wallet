package api

import (
	"e-wallet-wallet/helpers"
	"e-wallet-wallet/internal/interfaces"
	"e-wallet-wallet/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WalletHandler struct {
	WalletService interfaces.IWalletService
}

func (api *WalletHandler) Create(c *gin.Context) {
	var (
		log = helpers.Logger
		req *models.Wallet
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("failed to parse request body", err)
		helpers.SendResponseHTTP(
			c,
			http.StatusBadRequest,
			false,
			"failed to parse request body",
			nil,
		)
		c.Abort()
		return
	}

	if req.UserID == 0 {
		log.Error("user id is required")
		helpers.SendResponseHTTP(
			c,
			http.StatusBadRequest,
			false,
			"user id is required",
			nil,
		)
		c.Abort()
		return
	}

	err := api.WalletService.Create(c.Request.Context(), req)
	if err != nil {
		log.Error("failed to create wallet", err)
		helpers.SendResponseHTTP(
			c,
			http.StatusInternalServerError,
			false,
			"failed to create wallet",
			nil,
		)
		c.Abort()
		return
	}

	helpers.SendResponseHTTP(
		c,
		http.StatusOK,
		true,
		"wallet created successfully",
		req,
	)
}
