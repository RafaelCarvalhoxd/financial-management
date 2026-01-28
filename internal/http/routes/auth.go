package routes

import (
	"github.com/RafaelCarvalhoxd/financial-mangement/internal/apps/auth"
	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(r *gin.RouterGroup, handler *auth.Handler) {
	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)
}
