package transaction

import (
	"context"
	"math"
	"time"

	apperrors "github.com/RafaelCarvalhoxd/financial-management/internal/infra/errors"
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Create(ctx context.Context, description string, amount float64, date time.Time, transactionType string, categoryID int, userID int) (*Transaction, error) {

	normalized := math.Abs(amount)
	if transactionType == "expense" {
		normalized = -normalized
	}

	return s.repository.Create(ctx, description, normalized, date, transactionType, categoryID, userID)
}

func (s *Service) FindByID(ctx context.Context, id int, userID int) (*Transaction, error) {
	transaction, err := s.repository.FindByID(ctx, id, userID)
	if err != nil {
		return nil, err
	}
	if transaction == nil {
		return nil, apperrors.ErrNotFound
	}
	return transaction, nil
}

func (s *Service) FindAllGroupedByCategory(ctx context.Context, userID int, year *int, month *int) (*GroupedTransactionsModel, error) {
	transactions, err := s.repository.FindAll(ctx, userID, year, month)
	if err != nil {
		return nil, err
	}

	categoryMap := make(map[int]*CategoryGroupModel)
	var totalBalance, totalIncome, totalExpense float64

	for _, transaction := range transactions {
		totalBalance += transaction.Amount

		switch transaction.Type {
		case "income":
			totalIncome += transaction.Amount
		case "expense":
			totalExpense += math.Abs(transaction.Amount)
		}

		if categoryMap[transaction.CategoryID] == nil {
			categoryMap[transaction.CategoryID] = &CategoryGroupModel{
				CategoryID:   transaction.CategoryID,
				CategoryName: transaction.CategoryName,
				Balance:      0,
				Transactions: []*Transaction{},
			}
		}

		group := categoryMap[transaction.CategoryID]
		group.Balance += transaction.Amount
		group.Transactions = append(group.Transactions, transaction)
	}

	categories := make([]*CategoryGroupModel, 0, len(categoryMap))
	for _, group := range categoryMap {
		categories = append(categories, group)
	}

	return &GroupedTransactionsModel{
		TotalBalance: totalBalance,
		TotalIncome:  totalIncome,
		TotalExpense: totalExpense,
		Categories:   categories,
	}, nil
}

func (s *Service) Update(
	ctx context.Context,
	id int,
	userID int,
	description *string,
	amount *float64,
	date *time.Time,
	transactionType *string,
	categoryID *int,
) (*Transaction, error) {
	transaction, err := s.repository.FindByID(ctx, id, userID)
	if err != nil {
		return nil, err
	}
	if transaction == nil {
		return nil, apperrors.ErrNotFound
	}

	if amount != nil {
		normalized := math.Abs(*amount)
		if transaction.Type == "expense" {
			normalized = -normalized
		}
		*amount = normalized
	} else if transactionType != nil {
		normalized := math.Abs(transaction.Amount)
		if transaction.Type == "expense" {
			normalized = -normalized
		}
		amount = &normalized
	}

	return s.repository.Update(ctx, id, userID, description, amount, date, transactionType, categoryID)
}

func (s *Service) Delete(ctx context.Context, id int, userID int) error {
	transaction, err := s.repository.FindByID(ctx, id, userID)
	if err != nil {
		return err
	}
	if transaction == nil {
		return apperrors.ErrNotFound
	}
	return s.repository.Delete(ctx, id, userID)
}
