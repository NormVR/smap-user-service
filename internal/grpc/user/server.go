package user

import (
	"context"
	"errors"
	"log"
	"user-service/internal/domain/models"

	user_service "github.com/NormVR/smap_protobuf/gen/services/user_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	domain_errors "user-service/internal/domain/errors"
)

type UserService interface {
	GetUser(ctx context.Context, id int64) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
}

type GrpcSrv struct {
	user_service.UnimplementedUserServiceServer
	user UserService
}

func Register(grpcServer *grpc.Server, user UserService) {
	user_service.RegisterUserServiceServer(grpcServer, &GrpcSrv{
		user: user,
	})
}

func (s *GrpcSrv) GetUser(ctx context.Context, req *user_service.GetUserRequest) (*user_service.GetUserResponse, error) {
	userData, err := s.user.GetUser(ctx, req.UserId)

	if err != nil {
		log.Printf("Error getting user: %v", err)
		if errors.Is(err, domain_errors.ErrUserNotFound) {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}

		return nil, status.Errorf(codes.Internal, "internal error")
	}

	return &user_service.GetUserResponse{
		UserId:    userData.Id,
		Email:     userData.Email,
		Username:  userData.Username,
		Firstname: userData.Firstname,
		Lastname:  userData.Lastname,
	}, nil
}

func (s *GrpcSrv) UpdateUser(ctx context.Context, req *user_service.UpdateUserRequest) (*emptypb.Empty, error) {
	userData := &models.User{
		Id:        req.UserId,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
	}

	err := s.user.UpdateUser(ctx, userData)

	if err != nil {
		log.Printf("Error getting user: %v", err)
		if errors.Is(err, domain_errors.ErrUserNotFound) {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}

		return nil, status.Errorf(codes.Internal, "internal error")
	}

	return &emptypb.Empty{}, nil
}
