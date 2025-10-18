package app

import (
	grpcApp "user-service/internal/app/grpc"
	"user-service/internal/config"
	userService "user-service/internal/services/user"
	"user-service/internal/storage/postgres"
)

type App struct {
	GrpcSrv *grpcApp.App
}

func New(config *config.Config) *App {
	dbStorage, err := postgres.New(config.PostgresDsn)
	if err != nil {
		panic(err)
	}

	service := userService.New(dbStorage)
	app := grpcApp.New(service, config.GrpcPort)
	return &App{
		GrpcSrv: app,
	}
}
