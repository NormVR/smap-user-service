package app

import (
	grpcApp "user-service/internal/app/grpc"
	kafkaApp "user-service/internal/app/kafka"
	"user-service/internal/config"
	kafkaConsumer "user-service/internal/kafka"
	userService "user-service/internal/services/user"
	"user-service/internal/storage/postgres"
)

type App struct {
	GrpcSrv       *grpcApp.App
	KafkaConsumer *kafkaApp.App
}

func New(config *config.Config) *App {
	dbStorage, err := postgres.New(config.PostgresDsn)
	if err != nil {
		panic(err)
	}

	service := userService.New(dbStorage)
	grpcApp := grpcApp.New(service, config.GrpcPort)

	kafkaBrokers := []string{config.KafkaBrokers}
	kafkaConsumer := kafkaConsumer.New(kafkaBrokers, service)
	kafkaApp := kafkaApp.New(kafkaConsumer)

	return &App{
		GrpcSrv:       grpcApp,
		KafkaConsumer: kafkaApp,
	}
}
