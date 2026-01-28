package auth

import (
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

func (h *Handler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Failed to decode request"})
		return
	}

	ctx := c.Request.Context()
	user, err := h.service.Register(ctx, req.Name, req.Email, req.Password)
	if err != nil {
		helpers.HandleError(c, err)
		return
	}

	response := RegisterResponse{
		ID:    user.ID,
		Email: user.Email,
	}

	c.JSON(201, response)
}

func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Failed to decode request"})
		return
	}

	ctx := c.Request.Context()
	token, err := h.service.Login(ctx, req.Email, req.Password)
	if err != nil {
		helpers.HandleError(c, err)
		return
	}

	response := LoginResponse{
		Token: token,
	}

	c.JSON(200, response)
}