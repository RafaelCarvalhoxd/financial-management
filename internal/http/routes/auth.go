package routes

import (
	"github.com/RafaelCarvalhoxd/financial-management/internal/auth"
	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(r *gin.RouterGroup, handler *auth.Handler) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", handler.Register)
		auth.POST("/login", handler.Login)
	}
}
