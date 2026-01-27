package auth

import (
	"context"

	"github.com/RafaelCarvalhoxd/financial-mangement/internal/errors"
	"github.com/RafaelCarvalhoxd/financial-mangement/internal/user"
)

type Service struct {
	userRepository *user.Repository
}

func NewService(userRepository *user.Repository) *Service {
	return &Service{userRepository: userRepository}
}

func (s *Service) Register(ctx context.Context, name, email, password string) (*user.User, error) {
	user, err := s.userRepository.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return nil, errors.ErrConflict
	}
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return nil, err
	}
	return s.userRepository.Create(ctx, name, email, hashedPassword)
}

func (s *Service) Login(ctx context.Context, email, password, jwtSecret string) (string, error) {
	user, err := s.userRepository.FindByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.ErrNotFound
	}
	if !VerifyPassword(password, user.Password) {
		return "", errors.ErrUnauthorized
	}
	token, err := GenerateToken(user.ID, user.Email, jwtSecret)
	if err != nil {
		return "", err
	}
	return token, nil
}
