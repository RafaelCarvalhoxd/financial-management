package category

import (
	"strconv"

	"github.com/RafaelCarvalhoxd/financial-management/internal/http/helpers"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) CreateCategory(c *gin.Context) {
	var req CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Failed to decode request"})
		return
	}

	ctx := c.Request.Context()
	category, err := h.service.Create(ctx, req.Name, 1)
	if err != nil {
		helpers.HandleError(c, err)
		return
	}

	c.JSON(201, CategoryResponse{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	})
}

func (h *Handler) UpdateCategory(c *gin.Context) {
	var req UpdateCategoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Failed to decode request"})
		return
	}
	id := c.Param("id")
	categoryID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	ctx := c.Request.Context()
	category, err := h.service.Update(ctx, categoryID, req.Name, 1)
	if err != nil {
		helpers.HandleError(c, err)
		return
	}

	c.JSON(200, CategoryResponse{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	})
}

func (h *Handler) GetCategories(c *gin.Context) {
	ctx := c.Request.Context()
	categories, err := h.service.FindAll(ctx, 1)
	if err != nil {
		helpers.HandleError(c, err)
		return
	}

	responses := []CategoryResponse{}
	for _, category := range categories {
		responses = append(responses, CategoryResponse{
			ID:        category.ID,
			Name:      category.Name,
			CreatedAt: category.CreatedAt,
			UpdatedAt: category.UpdatedAt,
		})
	}

	c.JSON(200, responses)
}

func (h *Handler) GetCategory(c *gin.Context) {
	id := c.Param("id")
	ctx := c.Request.Context()
	categoryID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	category, err := h.service.FindByID(ctx, categoryID, 1)
	if err != nil {
		helpers.HandleError(c, err)
		return
	}

	c.JSON(200, CategoryResponse{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	})
}

func (h *Handler) DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	ctx := c.Request.Context()
	categoryID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	err = h.service.Delete(ctx, categoryID, 1)
	if err != nil {
		helpers.HandleError(c, err)
		return
	}

	c.JSON(200, gin.H{"message": "Category deleted successfully"})
}
