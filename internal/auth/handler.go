package auth

import (
	"github.com/RafaelCarvalhoxd/financial-mangement/internal/config"
	"github.com/RafaelCarvalhoxd/financial-mangement/internal/http/helpers"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
	config  *config.Config
}

func NewHandler(service *Service, cfg *config.Config) *Handler {
	return &Handler{
		service: service,
		config:  cfg,
	}
}

func (h *Handler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Erro ao decodificar requisição"})
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
		c.JSON(400, gin.H{"error": "Erro ao decodificar requisição"})
		return
	}

	ctx := c.Request.Context()
	token, err := h.service.Login(ctx, req.Email, req.Password, h.config.JWTSecret)
	if err != nil {
		helpers.HandleError(c, err)
		return
	}

	response := LoginResponse{
		Token: token,
	}

	c.JSON(200, response)
}
