package user

import (
	"context"
	"errors"
	"log"
	"time"

	apperrors "github.com/RafaelCarvalhoxd/financial-mangement/internal/errors"
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

func (r *Repository) Create(ctx context.Context, name, email, password string) (*User, error) {
	query := `
		INSERT INTO users (name, email, password, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, name, email, password, created_at, updated_at
	`

	now := time.Now()
	var user User

	err := r.db.QueryRow(
		ctx,
		query,
		name,
		email,
		password,
		now,
		now,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		log.Printf("Erro ao criar usuário no banco de dados: %v", err)
		return nil, apperrors.ErrInternalServerError
	}

	return &user, nil
}

func (r *Repository) FindByEmail(ctx context.Context, email string) (*User, error) {
	query := `
		SELECT id, name, email, password, created_at, updated_at FROM users WHERE email = $1
	`

	var user User

	err := r.db.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		log.Printf("Erro ao buscar usuário no banco de dados: %v", err)
		return nil, apperrors.ErrInternalServerError
	}

	return &user, nil
}
