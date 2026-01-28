package server

import (
	"github.com/RafaelCarvalhoxd/financial-mangement/internal/apps/auth"
	"github.com/RafaelCarvalhoxd/financial-mangement/internal/apps/category"
	"github.com/RafaelCarvalhoxd/financial-mangement/internal/http/routes"
	"github.com/gin-gonic/gin"
)

type Dependencies struct {
	AuthHandler     *auth.Handler
	CategoryHandler *category.Handler
}

func Config(deps *Dependencies) *gin.Engine {
	r := gin.Default()

	routes.SetupHealthRoutes(r)

	api := r.Group("/api")
	{
		routes.SetupAuthRoutes(api, deps.AuthHandler)
		routes.SetupCategoryRoutes(api, deps.CategoryHandler)
	}

	return r
}
