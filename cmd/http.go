package cmd

import (
	"e-wallet-wallet/external"
	"e-wallet-wallet/helpers"
	"e-wallet-wallet/internal/api"
	"e-wallet-wallet/internal/interfaces"
	"e-wallet-wallet/internal/repository"
	"e-wallet-wallet/internal/services"
	"github.com/gin-gonic/gin"
	"log"
)

func ServeHTTP() {
	d := dependencyInject()

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	err := r.SetTrustedProxies(nil)
	if err != nil {
		log.Fatal("Failed to set trusted proxies", err)
	}

	walletV1 := r.Group("/wallet/v1")
	walletV1.POST("/", d.WalletAPI.Create)
	walletV1.PUT("/balance/credit", d.MiddlewareValidateToken, d.WalletAPI.CreditBalance)
	walletV1.PUT("/balance/debit", d.MiddlewareValidateToken, d.WalletAPI.DebitBalance)
	walletV1.GET("/balance", d.MiddlewareValidateToken, d.WalletAPI.GetBalance)
	walletV1.GET("/history", d.MiddlewareValidateToken, d.WalletAPI.GetWalletHistory)

	exWalletV1 := walletV1.Group("/ex")
	exWalletV1.Use(d.MiddlewareSignatureValidation)
	exWalletV1.POST("/link", d.WalletAPI.CrateWalletLink)
	exWalletV1.PUT("/link/:wallet_id/confirmation", d.WalletAPI.WalletLinkConfirmation)
	exWalletV1.DELETE("/:wallet_id/unlink", d.WalletAPI.WalletUnlink)
	exWalletV1.GET("/:wallet_id/balance", d.WalletAPI.ExGetBalance)

	err = r.Run(":" + helpers.GetEnv("PORT", "8080"))
	if err != nil {
		log.Fatal("Failed to start server", err)
	}
	log.Println("Server started")
}

type Dependency struct {
	WalletAPI interfaces.IWalletHandler
	External  interfaces.IExternal
}

func dependencyInject() *Dependency {
	walletRepo := &repository.WalletRepository{DB: helpers.DB}
	walletSvc := &services.WalletService{WalletRepository: walletRepo}
	walletApi := &api.WalletHandler{WalletService: walletSvc}

	ext := &external.External{}

	return &Dependency{
		WalletAPI: walletApi,
		External:  ext,
	}
}
