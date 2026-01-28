package transaction

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	apperrors "github.com/RafaelCarvalhoxd/financial-management/internal/infra/errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) FindByID(ctx context.Context, id int, userID int) (*Transaction, error) {
	query := `
		SELECT t.id, t.description, t.amount, t.date, t.type, t.category_id, c.name AS category_name, t.user_id, t.created_at, t.updated_at
		FROM transactions t
		LEFT JOIN categories c ON t.category_id = c.id
		WHERE t.id = $1 AND t.user_id = $2
	`

	var transaction Transaction

	err := r.db.QueryRow(ctx, query, id, userID).Scan(
		&transaction.ID,
		&transaction.Description,
		&transaction.Amount,
		&transaction.Date,
		&transaction.Type,
		&transaction.CategoryID,
		&transaction.CategoryName,
		&transaction.UserID,
		&transaction.CreatedAt,
		&transaction.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		log.Printf("Erro ao buscar transação no banco de dados: %v", err)
		return nil, apperrors.ErrInternalServerError
	}

	return &transaction, nil
}

func (r *Repository) Create(ctx context.Context, description string, amount float64, date time.Time, transactionType string, categoryID int, userID int) (*Transaction, error) {
	query := `
		INSERT INTO transactions (description, amount, date, type, category_id, user_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
	`

	now := time.Now()
	var id int

	err := r.db.QueryRow(
		ctx,
		query,
		description,
		amount,
		date,
		transactionType,
		categoryID,
		userID,
		now,
		now,
	).Scan(&id)

	if err != nil {
		log.Printf("Erro ao criar transação no banco de dados: %v", err)
		return nil, apperrors.ErrInternalServerError
	}

	return r.FindByID(ctx, id, userID)
}

func (r *Repository) FindAll(ctx context.Context, userID int, year *int, month *int) ([]*Transaction, error) {
	query := `
		SELECT t.id, t.description, t.amount, t.date, t.type, t.category_id, c.name AS category_name, t.user_id, t.created_at, t.updated_at
		FROM transactions t
		LEFT JOIN categories c ON t.category_id = c.id
		WHERE t.user_id = $1
		  AND ($2::int IS NULL OR EXTRACT(YEAR FROM t.date)::int = $2::int)
		  AND ($3::int IS NULL OR EXTRACT(MONTH FROM t.date)::int = $3::int)
		ORDER BY t.date DESC, t.created_at DESC
	`

	var transactions []*Transaction

	rows, err := r.db.Query(ctx, query, userID, year, month)
	if err != nil {
		log.Printf("Erro ao buscar transações no banco de dados: %v", err)
		return nil, apperrors.ErrInternalServerError
	}
	defer rows.Close()

	for rows.Next() {
		var transaction Transaction
		err := rows.Scan(
			&transaction.ID,
			&transaction.Description,
			&transaction.Amount,
			&transaction.Date,
			&transaction.Type,
			&transaction.CategoryID,
			&transaction.CategoryName,
			&transaction.UserID,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
		)
		if err != nil {
			log.Printf("Erro ao scanear transação: %v", err)
			return nil, apperrors.ErrInternalServerError
		}
		transactions = append(transactions, &transaction)
	}

	return transactions, nil
}

func (r *Repository) Update(ctx context.Context, id int, userID int, description *string, amount *float64, date *time.Time, transactionType *string, categoryID *int) (*Transaction, error) {
	query := `UPDATE transactions SET updated_at = $1`
	args := []interface{}{time.Now()}
	argPos := 2

	if description != nil {
		query += `, description = $` + fmt.Sprintf("%d", argPos)
		args = append(args, *description)
		argPos++
	}

	if amount != nil {
		query += `, amount = $` + fmt.Sprintf("%d", argPos)
		args = append(args, *amount)
		argPos++
	}

	if date != nil {
		query += `, date = $` + fmt.Sprintf("%d", argPos)
		args = append(args, *date)
		argPos++
	}

	if transactionType != nil {
		query += `, type = $` + fmt.Sprintf("%d", argPos)
		args = append(args, *transactionType)
		argPos++
	}

	if categoryID != nil {
		query += `, category_id = $` + fmt.Sprintf("%d", argPos)
		args = append(args, *categoryID)
		argPos++
	}

	query += ` WHERE id = $` + fmt.Sprintf("%d", argPos) + ` AND user_id = $` + fmt.Sprintf("%d", argPos+1)
	args = append(args, id, userID)

	_, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		log.Printf("Erro ao atualizar transação no banco de dados: %v", err)
		return nil, apperrors.ErrInternalServerError
	}

	return r.FindByID(ctx, id, userID)
}

func (r *Repository) Delete(ctx context.Context, id int, userID int) error {
	query := `
		DELETE FROM transactions WHERE id = $1 AND user_id = $2
	`

	_, err := r.db.Exec(ctx, query, id, userID)
	if err != nil {
		log.Printf("Erro ao deletar transação no banco de dados: %v", err)
		return apperrors.ErrInternalServerError
	}

	return nil
}
