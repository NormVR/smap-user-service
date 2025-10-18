package user

import (
	"context"
	"user-service/internal/domain/models"
)

type UserService struct {
	dbStorage dbStorage
}

type dbStorage interface {
	GetUser(ctx context.Context, id int64) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
}

func New(dbStorage dbStorage) *UserService {
	return &UserService{
		dbStorage: dbStorage,
	}
}

func (s *UserService) GetUser(ctx context.Context, id int64) (*models.User, error) {
	userData, err := s.dbStorage.GetUser(ctx, id)

	if err != nil {
		return nil, err
	}

	return &models.User{
		Id:        userData.Id,
		Email:     userData.Email,
		Username:  userData.Username,
		Firstname: userData.Firstname,
		Lastname:  userData.Lastname,
	}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, user *models.User) error {
	err := s.dbStorage.UpdateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
