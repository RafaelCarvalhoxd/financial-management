package transaction

import (
	"time"
)

type Transaction struct {
	ID           int       `db:"id" json:"id"`
	Description  string    `db:"description" json:"description"`
	Amount       float64   `db:"amount" json:"amount"`
	Date         time.Time `db:"date" json:"date"`
	Type         string    `db:"type" json:"type"`
	CategoryID   int       `db:"category_id" json:"category_id"`
	CategoryName string    `db:"category_name" json:"category_name,omitempty"`
	UserID       int       `db:"user_id" json:"user_id"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}

type CategoryGroupModel struct {
	CategoryID   int
	CategoryName string
	Balance      float64
	Transactions []*Transaction
}

type GroupedTransactionsModel struct {
	TotalBalance float64
	TotalIncome  float64
	TotalExpense float64
	Categories   []*CategoryGroupModel
}
