package app

import (
	grpcApp "user-service/internal/app/grpc"
	"user-service/internal/config"
	userService "user-service/internal/services/user"
)

type App struct {
	GrpcSrv *grpcApp.App
}

func New(config *config.Config) *App {
	service := userService.New()
	app := grpcApp.New(service, config.GrpcPort)
	return &App{
		GrpcSrv: app,
	}
}
