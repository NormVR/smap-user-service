package kafka

import (
	"context"
	"user-service/internal/kafka"
)

type App struct {
	consumer *kafka.Consumer
}

func New(consumer *kafka.Consumer) *App {
	return &App{
		consumer: consumer,
	}
}

func (a *App) Listen(ctx context.Context) {
	a.consumer.Listen(ctx)
}

func (a *App) Stop() {
	a.consumer.Stop()
}
