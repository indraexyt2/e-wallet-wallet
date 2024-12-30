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

func (api *WalletHandler) CreditBalance(c *gin.Context) {
	var (
		log = helpers.Logger
		req *models.TransactionRequest
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

	if req.Validate() != nil {
		log.Error("failed to validate request body")
		helpers.SendResponseHTTP(
			c,
			http.StatusBadRequest,
			false,
			"data validation failed",
			nil,
		)
		return
	}

	token, ok := c.Get("token")
	if !ok {
		log.Error("token is required")
		helpers.SendResponseHTTP(
			c,
			http.StatusBadRequest,
			false,
			"token is required",
			nil,
		)
		return
	}

	tokenData, ok := token.(*models.TokenData)
	if !ok {
		log.Error("token is required")
		helpers.SendResponseHTTP(
			c,
			http.StatusBadRequest,
			false,
			"token is required",
			nil,
		)
		return
	}

	resp, err := api.WalletService.CreditBalance(c.Request.Context(), int(tokenData.UserID), req)
	if err != nil {
		log.Error("failed to credit balance", err)
		helpers.SendResponseHTTP(
			c,
			http.StatusInternalServerError,
			false,
			"failed to credit balance",
			nil,
		)
		return
	}

	helpers.SendResponseHTTP(
		c,
		http.StatusOK,
		true,
		"success",
		resp,
	)
}

func (api *WalletHandler) DebitBalance(c *gin.Context) {
	var (
		log = helpers.Logger
		req *models.TransactionRequest
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

	if req.Validate() != nil {
		log.Error("failed to validate request body")
		helpers.SendResponseHTTP(
			c,
			http.StatusBadRequest,
			false,
			"data validation failed",
			nil,
		)
		return
	}

	token, ok := c.Get("token")
	if !ok {
		log.Error("token is required")
		helpers.SendResponseHTTP(
			c,
			http.StatusBadRequest,
			false,
			"token is required",
			nil,
		)
		return
	}

	tokenData, ok := token.(*models.TokenData)
	if !ok {
		log.Error("token is required")
		helpers.SendResponseHTTP(
			c,
			http.StatusBadRequest,
			false,
			"token is required",
			nil,
		)
		return
	}

	resp, err := api.WalletService.DebitBalance(c.Request.Context(), int(tokenData.UserID), req)
	if err != nil {
		log.Error("failed to credit balance", err)
		helpers.SendResponseHTTP(
			c,
			http.StatusInternalServerError,
			false,
			"failed to debit balance",
			nil,
		)
		return
	}

	helpers.SendResponseHTTP(
		c,
		http.StatusOK,
		true,
		"success",
		resp,
	)
}

func (api *WalletHandler) GetBalance(c *gin.Context) {
	var (
		log = helpers.Logger
	)

	token, ok := c.Get("token")
	if !ok {
		log.Error("token is required")
		helpers.SendResponseHTTP(
			c,
			http.StatusBadRequest,
			false,
			"token is required",
			nil,
		)
		return
	}

	tokenData, ok := token.(*models.TokenData)
	if !ok {
		log.Error("token is required")
		helpers.SendResponseHTTP(
			c,
			http.StatusBadRequest,
			false,
			"token is required",
			nil,
		)
		return
	}

	resp, err := api.WalletService.GetBalance(c.Request.Context(), int(tokenData.UserID))
	if err != nil {
		log.Error("failed to get balance", err)
		helpers.SendResponseHTTP(
			c,
			http.StatusInternalServerError,
			false,
			"failed to get balance",
			nil,
		)
		return
	}

	helpers.SendResponseHTTP(
		c,
		http.StatusOK,
		true,
		"success",
		resp,
	)
}

func (api *WalletHandler) GetWalletHistory(c *gin.Context) {
	var (
		log   = helpers.Logger
		param models.WalletHistoryParam
	)

	if err := c.ShouldBindQuery(&param); err != nil {
		log.Error("failed to parse request body: ", err)
		helpers.SendResponseHTTP(
			c,
			http.StatusBadRequest,
			false,
			"bad request",
			nil,
		)
		return
	}

	if param.WalletTransactionType != "" {
		if param.WalletTransactionType != "CREDIT" && param.WalletTransactionType != "DEBIT" {
			log.Error("invalid wallet transaction type")
			helpers.SendResponseHTTP(
				c,
				http.StatusBadRequest,
				false,
				"invalid wallet transaction type",
				nil,
			)
			return
		}
	}

	token, ok := c.Get("token")
	if !ok {
		log.Error("token is required")
		helpers.SendResponseHTTP(
			c,
			http.StatusBadRequest,
			false,
			"token is required",
			nil,
		)
		return
	}

	tokenData, ok := token.(*models.TokenData)
	if !ok {
		log.Error("token is required")
		helpers.SendResponseHTTP(
			c,
			http.StatusBadRequest,
			false,
			"token is required",
			nil,
		)
		return
	}

	resp, err := api.WalletService.GetWalletHistory(c.Request.Context(), int(tokenData.UserID), param)
	if err != nil {
		log.Error("failed to get wallet history", err)
		helpers.SendResponseHTTP(
			c,
			http.StatusInternalServerError,
			false,
			"failed to get wallet history",
			nil,
		)
		return
	}

	helpers.SendResponseHTTP(
		c,
		http.StatusOK,
		true,
		"success",
		resp,
	)
}
