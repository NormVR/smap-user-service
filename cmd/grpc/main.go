package main

import (
	"log"
	"os"
	"os/signal"
	"user-service/internal/app"
	"user-service/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatal("failed to load config:", err)
	}

	application := app.New(cfg)
	go application.GrpcSrv.Serve()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	application.GrpcSrv.Stop()
}
