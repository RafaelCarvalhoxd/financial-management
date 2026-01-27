package server

import (
	"github.com/RafaelCarvalhoxd/financial-mangement/internal/auth"
	"github.com/RafaelCarvalhoxd/financial-mangement/internal/http/routes"
	"github.com/gin-gonic/gin"
)

type Dependencies struct {
	AuthHandler *auth.Handler
}

func Config(deps *Dependencies) *gin.Engine {
	r := gin.Default()

	routes.SetupHealthRoutes(r)

	api := r.Group("/api")
	{
		routes.SetupAuthRoutes(api, deps.AuthHandler)

	}

	return r
}
