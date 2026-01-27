package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupHealthRoutes(r *gin.Engine) {
	r.GET("/health", handleHealth)
}

func handleHealth(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "ok",
		"service": "financial-management-api",
	})
}
