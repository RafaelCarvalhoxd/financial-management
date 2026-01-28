package auth

import (
	"context"

	"github.com/RafaelCarvalhoxd/financial-management/internal/user"
	"github.com/RafaelCarvalhoxd/financial-management/internal/infra/errors"
)

type Service struct {
	userRepository *user.Repository
	jwtSecret      string
}

func NewService(userRepository *user.Repository, jwtSecret string) *Service {
	return &Service{
		userRepository: userRepository,
		jwtSecret:      jwtSecret,
	}
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

func (s *Service) Login(ctx context.Context, email, password string) (string, error) {
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
	token, err := GenerateToken(user.ID, user.Email, s.jwtSecret)
	if err != nil {
		return "", err
	}
	return token, nil
}