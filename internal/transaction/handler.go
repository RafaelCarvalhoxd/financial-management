package transaction

import (
	"strconv"
	"time"

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

func (h *Handler) CreateTransaction(c *gin.Context) {
	var request CreateTransactionRequest
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Erro ao decodificar requisição"})
		return
	}

	ctx := c.Request.Context()
	transaction, err := h.service.Create(ctx, request.Description, request.Amount, request.Date, request.Type, request.CategoryID, 1)
	if err != nil {
		helpers.HandleError(c, err)
		return
	}

	c.JSON(201, TransactionResponse{
		ID:          transaction.ID,
		Description: transaction.Description,
		Amount:      transaction.Amount,
		Date:        transaction.Date,
		Type:        transaction.Type,
		UserID:      transaction.UserID,
		CreatedAt:   transaction.CreatedAt,
		UpdatedAt:   transaction.UpdatedAt,
	})
}

func (h *Handler) GetTransaction(c *gin.Context) {
	id := c.Param("id")
	ctx := c.Request.Context()
	transactionID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID inválido"})
		return
	}
	transaction, err := h.service.FindByID(ctx, transactionID, 1)
	if err != nil {
		helpers.HandleError(c, err)
		return
	}

	c.JSON(200, TransactionResponse{
		ID:          transaction.ID,
		Description: transaction.Description,
		Amount:      transaction.Amount,
		Date:        transaction.Date,
		Type:        transaction.Type,
		UserID:      transaction.UserID,
		CreatedAt:   transaction.CreatedAt,
		UpdatedAt:   transaction.UpdatedAt,
	})
}

func (h *Handler) GetTransactions(c *gin.Context) {
	ctx := c.Request.Context()

	var year, month *int
	if yearStr := c.Query("year"); yearStr != "" {
		y, err := strconv.Atoi(yearStr)
		if err != nil || y <= 0 {
			c.JSON(400, gin.H{"error": "Parâmetro 'year' inválido"})
			return
		}
		year = &y
	}
	if monthStr := c.Query("month"); monthStr != "" {
		m, err := strconv.Atoi(monthStr)
		if err != nil || m < 1 || m > 12 {
			c.JSON(400, gin.H{"error": "Parâmetro 'month' inválido (1-12)"})
			return
		}
		month = &m
	}

	if (year != nil && month == nil) || (year == nil && month != nil) {
		c.JSON(400, gin.H{"error": "Os parâmetros 'year' e 'month' devem ser fornecidos juntos"})
		return
	}

	if year == nil && month == nil {
		now := time.Now()
		y := now.Year()
		m := int(now.Month())
		year = &y
		month = &m
	}

	grouped, err := h.service.FindAllGroupedByCategory(ctx, 1, year, month)
	if err != nil {
		helpers.HandleError(c, err)
		return
	}

	response := GroupedTransactionsResponse{
		TotalBalance: grouped.TotalBalance,
		TotalIncome:  grouped.TotalIncome,
		TotalExpense: grouped.TotalExpense,
		Categories:   make([]CategoryGroup, 0, len(grouped.Categories)),
	}

	for _, category := range grouped.Categories {
		categoryGroup := CategoryGroup{
			CategoryID:   category.CategoryID,
			CategoryName: category.CategoryName,
			Balance:      category.Balance,
			Transactions: make([]TransactionResponse, 0, len(category.Transactions)),
		}

		for _, transaction := range category.Transactions {
			categoryGroup.Transactions = append(categoryGroup.Transactions, TransactionResponse{
				ID:          transaction.ID,
				Description: transaction.Description,
				Amount:      transaction.Amount,
				Date:        transaction.Date,
				Type:        transaction.Type,
				UserID:      transaction.UserID,
				CreatedAt:   transaction.CreatedAt,
				UpdatedAt:   transaction.UpdatedAt,
			})
		}

		response.Categories = append(response.Categories, categoryGroup)
	}

	c.JSON(200, response)
}

func (h *Handler) UpdateTransaction(c *gin.Context) {
	id := c.Param("id")
	ctx := c.Request.Context()
	transactionID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID inválido"})
		return
	}
	var request UpdateTransactionRequest
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Erro ao decodificar requisição"})
		return
	}
	transaction, err := h.service.Update(ctx, transactionID, 1, request.Description, request.Amount, request.Date, request.Type, request.CategoryID)
	if err != nil {
		helpers.HandleError(c, err)
		return
	}
	c.JSON(200, TransactionResponse{
		ID:          transaction.ID,
		Description: transaction.Description,
		Amount:      transaction.Amount,
		Date:        transaction.Date,
		Type:        transaction.Type,
		UserID:      transaction.UserID,
		CreatedAt:   transaction.CreatedAt,
		UpdatedAt:   transaction.UpdatedAt,
	})
}

func (h *Handler) DeleteTransaction(c *gin.Context) {
	id := c.Param("id")
	ctx := c.Request.Context()
	transactionID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID inválido"})
		return
	}
	err = h.service.Delete(ctx, transactionID, 1)
	if err != nil {
		helpers.HandleError(c, err)
		return
	}
	c.JSON(200, gin.H{"message": "Transação deletada com sucesso"})
}
