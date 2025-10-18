package postgres

import (
	"context"
	"errors"
	"fmt"
	"log"
	domainErrors "user-service/internal/domain/errors"
	"user-service/internal/domain/models"

	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Storage struct {
	db *pgx.Conn
}

func New(storagePath string) (*Storage, error) {

	conn, err := pgx.Connect(context.Background(), storagePath)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &Storage{db: conn}, nil
}

func (s *Storage) Close(ctx context.Context) {
	s.db.Close(ctx)
}

func (s *Storage) GetUser(ctx context.Context, id int64) (*models.User, error) {
	var user models.User
	err := s.db.QueryRow(ctx, `SELECT id, email, username, firstname, lastname FROM users WHERE id=$1`, id).Scan(
		&user.Id,
		&user.Email,
		&user.Username,
		&user.Firstname,
		&user.Lastname,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domainErrors.ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &user, nil
}

func (s *Storage) UpdateUser(ctx context.Context, user *models.User) error {

	res, err := s.db.Exec(ctx,
		`UPDATE users SET firstname=$1, lastname=$2 WHERE id=$3`,
		user.Firstname,
		user.Lastname,
		user.Id,
	)

	if err != nil {
		return err
	}

	if res.RowsAffected() == 0 {
		return domainErrors.ErrUserNotFound
	}

	return nil
}
