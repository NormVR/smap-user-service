package user

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"
	"user-service/internal/domain/models"

	domain_errors "user-service/internal/domain/errors"

	user_service "github.com/NormVR/smap_protobuf/gen/services/user_service"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserService interface {
	GetUser(ctx context.Context, id uuid.UUID) (*models.User, error)
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

	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fmt.Errorf("User ID converting error: %v", err)
	}

	userData, err := s.user.GetUser(ctx, userId)

	if err != nil {
		log.Printf("Error getting user: %v", err)
		if errors.Is(err, domain_errors.ErrUserNotFound) {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}

		return nil, status.Errorf(codes.Internal, "internal error")
	}

	return &user_service.GetUserResponse{
		UserId:    userData.Id.String(),
		Username:  userData.Username,
		Firstname: userData.Firstname,
		Lastname:  userData.Lastname,
		Bio:       userData.Bio,
		AvatarUrl: userData.AvatarUrl,
		Website:   userData.Website,
		Location:  userData.Location,
		BirthDate: userData.BirthDate.Format(time.RFC3339),
		Gender:    userData.Gender,
		Telephone: userData.Telephone,
	}, nil
}

func (s *GrpcSrv) UpdateUser(ctx context.Context, req *user_service.UpdateUserRequest) (*emptypb.Empty, error) {

	birthDate, err := time.Parse(time.RFC3339, req.BirthDate)
	if err != nil {
		log.Printf("Error parsing birth date: %v", err)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		log.Printf("User ID converting error: %v", err)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	userData := &models.User{
		Id:        userId,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Bio:       req.Bio,
		AvatarUrl: req.AvatarUrl,
		Website:   req.Website,
		Location:  req.Location,
		BirthDate: birthDate,
		Gender:    req.Gender,
		Telephone: req.Telephone,
	}

	err = s.user.UpdateUser(ctx, userData)

	if err != nil {
		log.Printf("Error getting user: %v", err)
		if errors.Is(err, domain_errors.ErrUserNotFound) {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}

		return nil, status.Errorf(codes.Internal, "internal error")
	}

	return &emptypb.Empty{}, nil
}
