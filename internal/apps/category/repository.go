package category

import (
	"context"
	"errors"
	"log"
	"time"

	apperrors "github.com/RafaelCarvalhoxd/financial-management/internal/errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, name string, userID int) (*Category, error) {
	query := `
		INSERT INTO categories (name, user_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id, name, user_id, created_at, updated_at
	`

	now := time.Now()
	var category Category

	err := r.db.QueryRow(ctx, query, name, userID, now, now).Scan(
		&category.ID,
		&category.Name,
		&category.UserID,
		&category.CreatedAt,
		&category.UpdatedAt,
	)

	if err != nil {
		log.Printf("Erro ao criar categoria no banco de dados: %v", err)
		return nil, apperrors.ErrInternalServerError
	}

	return &category, nil
}

func (r *Repository) FindAll(ctx context.Context, userID int) ([]*Category, error) {
	query := `
		SELECT id, name, user_id, created_at, updated_at FROM categories WHERE user_id = $1
	`

	var categories []*Category

	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		log.Printf("Erro ao buscar categorias no banco de dados: %v", err)
		return nil, apperrors.ErrInternalServerError
	}
	defer rows.Close()

	for rows.Next() {
		var category Category
		err := rows.Scan(&category.ID, &category.Name, &category.UserID, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			log.Printf("Erro ao scanear categoria: %v", err)
			return nil, apperrors.ErrInternalServerError
		}
		categories = append(categories, &category)
	}

	return categories, nil
}

func (r *Repository) FindByID(ctx context.Context, id int, userID int) (*Category, error) {
	query := `
		SELECT id, name, user_id, created_at, updated_at FROM categories WHERE id = $1 AND user_id = $2
	`

	var category Category

	err := r.db.QueryRow(ctx, query, id, userID).Scan(&category.ID, &category.Name, &category.UserID, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		log.Printf("Erro ao buscar categoria por ID no banco de dados: %v", err)
		return nil, apperrors.ErrInternalServerError
	}

	return &category, nil
}

func (r *Repository) Update(ctx context.Context, id int, name string, userID int) (*Category, error) {
	query := `
		UPDATE categories SET name = $1, updated_at = $2 WHERE id = $3 AND user_id = $4
		RETURNING id, name, user_id, created_at, updated_at
	`

	now := time.Now()
	var category Category

	err := r.db.QueryRow(ctx, query, name, now, id, userID).Scan(&category.ID, &category.Name, &category.UserID, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		log.Printf("Erro ao atualizar categoria no banco de dados: %v", err)
		return nil, apperrors.ErrInternalServerError
	}

	return &category, nil
}

func (r *Repository) Delete(ctx context.Context, id int, userID int) error {
	query := `
		DELETE FROM categories WHERE id = $1 AND user_id = $2
	`

	_, err := r.db.Exec(ctx, query, id, userID)
	if err != nil {
		log.Printf("Erro ao deletar categoria no banco de dados: %v", err)
		return apperrors.ErrInternalServerError
	}

	return nil
}

func (r *Repository) FindByName(ctx context.Context, name string, userID int) (*Category, error) {
	query := `
		SELECT id, name, user_id, created_at, updated_at FROM categories WHERE name = $1 AND user_id = $2
	`

	var category Category

	err := r.db.QueryRow(ctx, query, name, userID).Scan(&category.ID, &category.Name, &category.UserID, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		log.Printf("Erro ao buscar categoria por nome no banco de dados: %v", err)
		return nil, apperrors.ErrInternalServerError
	}

	return &category, nil
}
