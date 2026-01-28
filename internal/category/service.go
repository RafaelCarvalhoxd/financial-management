package category

import (
	"context"

	apperrors "github.com/RafaelCarvalhoxd/financial-management/internal/infra/errors"
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) Create(ctx context.Context, name string, userID int) (*Category, error) {
	existingCategory, err := s.repository.FindByName(ctx, name, userID)
	if err != nil {
		return nil, err
	}
	if existingCategory != nil {
		return nil, apperrors.ErrConflict
	}
	category, err := s.repository.Create(ctx, name, userID)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (s *Service) FindAll(ctx context.Context, userID int) ([]*Category, error) {
	categories, err := s.repository.FindAll(ctx, userID)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (s *Service) FindByID(ctx context.Context, id int, userID int) (*Category, error) {
	category, err := s.repository.FindByID(ctx, id, userID)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, apperrors.ErrNotFound
	}
	return category, nil
}

func (s *Service) Update(ctx context.Context, id int, name string, userID int) (*Category, error) {
	existingCategory, err := s.repository.FindByID(ctx, id, userID)
	if err != nil {
		return nil, err
	}
	if existingCategory == nil {
		return nil, apperrors.ErrNotFound
	}
	existingCategoryName, err := s.repository.FindByName(ctx, name, userID)
	if err != nil {
		return nil, err
	}
	if existingCategoryName != nil && existingCategoryName.ID != id && existingCategoryName.Name == name {
		return nil, apperrors.ErrConflict
	}
	category, err := s.repository.Update(ctx, id, name, userID)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (s *Service) Delete(ctx context.Context, id int, userID int) error {
	existingCategory, err := s.repository.FindByID(ctx, id, userID)
	if err != nil {
		return err
	}
	if existingCategory == nil {
		return apperrors.ErrNotFound
	}
	err = s.repository.Delete(ctx, id, userID)
	if err != nil {
		return err
	}
	return nil
}