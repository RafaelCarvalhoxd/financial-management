package transaction

import "time"

type CreateTransactionRequest struct {
	Description string    `json:"description" binding:"required"`
	Amount      float64   `json:"amount" binding:"required"`
	Date        time.Time `json:"date" binding:"required"`
	Type        string    `json:"type" binding:"required,oneof=income expense"`
	CategoryID  int       `json:"category_id" binding:"required"`
}

type UpdateTransactionRequest struct {
	Description *string    `json:"description,omitempty"`
	Amount      *float64   `json:"amount,omitempty"`
	Date        *time.Time `json:"date,omitempty"`
	Type        *string    `json:"type,omitempty" binding:"omitempty,oneof=income expense"`
	CategoryID  *int       `json:"category_id,omitempty"`
}

type TransactionResponse struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Date        time.Time `json:"date"`
	Type        string    `json:"type"`
	UserID      int       `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CategoryGroup struct {
	CategoryID   int                   `json:"category_id"`
	CategoryName string                `json:"category_name"`
	Balance      float64               `json:"balance"`
	Transactions []TransactionResponse `json:"transactions"`
}

type GroupedTransactionsResponse struct {
	TotalBalance float64         `json:"total_balance"`
	TotalIncome  float64         `json:"total_income"`
	TotalExpense float64         `json:"total_expense"`
	Categories   []CategoryGroup `json:"categories"`
}
