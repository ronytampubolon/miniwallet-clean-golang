package routes

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/ronytampubolon/miniwallet/handlers/wallet"
	"github.com/ronytampubolon/miniwallet/middlewares"
	repositories "github.com/ronytampubolon/miniwallet/repositories/wallet"
	services "github.com/ronytampubolon/miniwallet/services/wallet"

	"gorm.io/gorm"
)

func InitRoute(db *gorm.DB, route *gin.Engine) {

	/**
	@register Init Wallet Handler, Repository & Service
	*/
	miniWalletRepo := repositories.NewWalletRepository(db)
	miniWalletService := services.NewInitWalletService(miniWalletRepo)
	miniWalletHandler := handlers.NewWalletHandler(miniWalletService)

	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/init", miniWalletHandler.InitWalletHandler)
	groupRouteWallet := route.Group("/api/v1/wallet").Use(middlewares.AuthenticatedUser())
	groupRouteWallet.GET("/", miniWalletHandler.GetBalanceHandler)
	groupRouteWallet.POST("/", miniWalletHandler.EnableWalletHandler)
	groupRouteWallet.PATCH("/", miniWalletHandler.DisableWalletHandler)
	groupRouteWallet.POST("/deposits", miniWalletHandler.DepositHandler)
	groupRouteWallet.POST("/withdrawals", miniWalletHandler.WithdrawnHandler)
}
