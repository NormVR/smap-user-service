package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"
	"user-service/internal/domain/models"
	userService "user-service/internal/services/user"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader      *kafka.Reader
	userService *userService.UserService
}

type UserCreatedEvent struct {
	UserID    uuid.UUID `json:"user_id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

func New(kafkaBrokers []string, userService *userService.UserService) *Consumer {
	return &Consumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:        kafkaBrokers,
			Topic:          "user-created",
			GroupID:        "smap-user-service",
			CommitInterval: time.Second,
		}),
		userService: userService,
	}
}

func (c *Consumer) Listen(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return c.reader.Close()
		default:
			msg, err := c.reader.FetchMessage(ctx)
			if err != nil {
				if ctx.Err() != nil {
					return ctx.Err()
				}
				log.Println(err)
				continue
			}

			if err := c.processMessage(ctx, msg); err != nil {
				log.Printf("error processing message: %v", err)
				continue
			}

			if err := c.reader.CommitMessages(ctx, msg); err != nil {
				log.Printf("error committing message: %v", err)
			}
		}
	}
}

func (c *Consumer) processMessage(ctx context.Context, msg kafka.Message) error {
	var event UserCreatedEvent

	if err := json.Unmarshal(msg.Value, &event); err != nil {
		return fmt.Errorf("unmarshal event error: %w", err)
	}

	log.Printf("Received message with offset %d\n", event.UserID)

	user := models.User{
		Id:       event.UserID,
		Username: event.Username,
	}

	err := c.userService.CreateUser(ctx, &user)
	if err != nil {
		return fmt.Errorf("create user error: %w", err)
	}

	return nil
}

func (c *Consumer) Stop() {
	c.reader.Close()
}
