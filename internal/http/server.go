package http

import (
	"github.com/RafaelCarvalhoxd/financial-management/internal/auth"
	"github.com/RafaelCarvalhoxd/financial-management/internal/category"
	"github.com/RafaelCarvalhoxd/financial-management/internal/http/routes"
	"github.com/RafaelCarvalhoxd/financial-management/internal/transaction"
	"github.com/gin-gonic/gin"
)

type Dependencies struct {
	AuthHandler        *auth.Handler
	CategoryHandler    *category.Handler
	TransactionHandler *transaction.Handler
}

func Config(deps *Dependencies) *gin.Engine {
	r := gin.Default()

	routes.SetupHealthRoutes(r)

	api := r.Group("/api")
	{
		routes.SetupAuthRoutes(api, deps.AuthHandler)
		routes.SetupCategoryRoutes(api, deps.CategoryHandler)
		routes.SetupTransactionRoutes(api, deps.TransactionHandler)
	}

	return r
}
