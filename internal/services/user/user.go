package user

import (
	"user-service/internal/models/user"
)

type UserService struct {
}

func New() *UserService {
	return &UserService{}

}

func (s *UserService) GetUser(id int64) (*user.User, error) {
	return &user.User{
		Id:        1,
		Email:     "test@test.com",
		Username:  "Test Username",
		Firstname: "Test Firstname",
		Lastname:  "Test Lastname",
	}, nil
}

func (s *UserService) UpdateUser(user *user.User) error {
	return nil
}
