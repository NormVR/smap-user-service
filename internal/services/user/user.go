package user

import (
	"context"
	"user-service/internal/domain/models"

	"github.com/google/uuid"
)

type UserService struct {
	dbStorage dbStorage
}

type dbStorage interface {
	GetUser(ctx context.Context, id uuid.UUID) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
	CreateUser(ctx context.Context, user *models.User) error
}

func New(dbStorage dbStorage) *UserService {
	return &UserService{
		dbStorage: dbStorage,
	}
}

func (s *UserService) GetUser(ctx context.Context, id uuid.UUID) (*models.User, error) {
	userData, err := s.dbStorage.GetUser(ctx, id)

	if err != nil {
		return nil, err
	}

	return userData, nil
}

func (s *UserService) UpdateUser(ctx context.Context, user *models.User) error {
	err := s.dbStorage.UpdateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) CreateUser(ctx context.Context, user *models.User) error {
	err := s.dbStorage.CreateUser(ctx, user)

	if err != nil {
		return err
	}

	return nil
}
