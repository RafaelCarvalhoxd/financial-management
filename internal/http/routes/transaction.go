package routes

import (
	"github.com/RafaelCarvalhoxd/financial-management/internal/transaction"
	"github.com/gin-gonic/gin"
)

func SetupTransactionRoutes(router *gin.RouterGroup, handler *transaction.Handler) {
	transaction := router.Group("/transactions")
	{
		transaction.POST("", handler.CreateTransaction)
		transaction.GET("", handler.GetTransactions)
		transaction.GET("/:id", handler.GetTransaction)
		transaction.PUT("/:id", handler.UpdateTransaction)
		transaction.DELETE("/:id", handler.DeleteTransaction)
	}
}
