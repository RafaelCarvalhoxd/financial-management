package routes

import (
	"github.com/RafaelCarvalhoxd/financial-management/internal/category"
	"github.com/gin-gonic/gin"
)

func SetupCategoryRoutes(router *gin.RouterGroup, handler *category.Handler) {
	category := router.Group("/categories")
	{
		category.POST("", handler.CreateCategory)
		category.GET("", handler.GetCategories)
		category.GET("/:id", handler.GetCategory)
		category.PUT("/:id", handler.UpdateCategory)
		category.DELETE("/:id", handler.DeleteCategory)
	}
}
