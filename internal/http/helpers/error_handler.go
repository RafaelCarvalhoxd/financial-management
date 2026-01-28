package helpers

import (
	"log"

	"github.com/RafaelCarvalhoxd/financial-management/internal/infra/errors"
	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
	log.Printf("Handler error: %v", err)

	switch err {
	case errors.ErrConflict:
		c.JSON(409, gin.H{"error": "Resource already exists"})
	case errors.ErrNotFound:
		c.JSON(404, gin.H{"error": "Resource not found"})
	case errors.ErrUnauthorized:
		c.JSON(401, gin.H{"error": "Invalid credentials"})
	case errors.ErrInvalidInput:
		c.JSON(400, gin.H{"error": "Invalid data"})
	case errors.ErrBadRequest:
		c.JSON(400, gin.H{"error": "Invalid request"})
	case errors.ErrForbidden:
		c.JSON(403, gin.H{"error": "Access denied"})
	case errors.ErrInternalServer, errors.ErrInternalServerError:
		c.JSON(500, gin.H{"error": "Internal server error"})
	case errors.ErrTooManyRequests:
		c.JSON(429, gin.H{"error": "Too many requests"})
	case errors.ErrUnprocessableEntity:
		c.JSON(422, gin.H{"error": "Unprocessable entity"})
	default:
		c.JSON(500, gin.H{"error": "Internal server error"})
	}
}
