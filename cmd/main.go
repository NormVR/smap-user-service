package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"user-service/internal/app"
	"user-service/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err != nil {
		log.Fatal("failed to load config:", err)
	}

	application := app.New(cfg)
	go application.GrpcSrv.Serve()
	go application.KafkaConsumer.Listen(ctx)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	application.GrpcSrv.Stop()
	application.KafkaConsumer.Stop()
}
