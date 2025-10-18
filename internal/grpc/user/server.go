package user

import (
	"context"

	userService "github.com/NormVR/smap_protobuf/gen/services/user_service"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"user-service/internal/models/user"
)

type UserService interface {
	GetUser(id int64) (*user.User, error)
	UpdateUser(user *user.User) error
}

type GrpcSrv struct {
	userService.UnimplementedUserServiceServer
	user UserService
}

func Register(grpcServer *grpc.Server, user UserService) {
	userService.RegisterUserServiceServer(grpcServer, &GrpcSrv{
		user: user,
	})
}

func (s *GrpcSrv) GetUser(ctx context.Context, req *userService.GetUserRequest) (*userService.GetUserResponse, error) {
	userData, err := s.user.GetUser(req.UserId)

	if err != nil {
		return nil, err
	}

	return &userService.GetUserResponse{
		UserId:    userData.Id,
		Email:     userData.Email,
		Username:  userData.Username,
		Firstname: userData.Firstname,
		Lastname:  userData.Lastname,
	}, nil
}

func (s *GrpcSrv) UpdateUser(ctx context.Context, req *userService.UpdateUserRequest) (*emptypb.Empty, error) {
	userData := &user.User{
		Id:        req.UserId,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
	}

	err := s.user.UpdateUser(userData)

	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
