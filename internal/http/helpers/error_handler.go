package helpers

import (
	"log"

	"github.com/RafaelCarvalhoxd/financial-mangement/internal/errors"
	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
	log.Printf("Erro no handler: %v", err)

	switch err {
	case errors.ErrConflict:
		c.JSON(409, gin.H{"error": "Email já cadastrado"})
	case errors.ErrNotFound:
		c.JSON(404, gin.H{"error": "Usuário não encontrado"})
	case errors.ErrUnauthorized:
		c.JSON(401, gin.H{"error": "Credenciais inválidas"})
	case errors.ErrInvalidInput:
		c.JSON(400, gin.H{"error": "Dados inválidos"})
	case errors.ErrBadRequest:
		c.JSON(400, gin.H{"error": "Requisição inválida"})
	case errors.ErrForbidden:
		c.JSON(403, gin.H{"error": "Acesso negado"})
	case errors.ErrInternalServer, errors.ErrInternalServerError:
		c.JSON(500, gin.H{"error": "Erro interno do servidor"})
	case errors.ErrTooManyRequests:
		c.JSON(429, gin.H{"error": "Muitas requisições"})
	case errors.ErrUnprocessableEntity:
		c.JSON(422, gin.H{"error": "Entidade não processável"})
	default:
		c.JSON(500, gin.H{"error": "Erro interno do servidor"})
	}
}
