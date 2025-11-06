package postgres

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	domainErrors "user-service/internal/domain/errors"
	"user-service/internal/domain/models"

	"github.com/google/uuid"
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

func (s *Storage) CreateUser(ctx context.Context, user *models.User) error {

	_, err := s.GetUser(ctx, user.Id)

	if err == nil {
		return nil
	}

	if !errors.Is(err, domainErrors.ErrUserNotFound) {
		return fmt.Errorf("could not load user to check existance before create: %w", err)
	}

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, `INSERT INTO user_profiles (id, username) VALUES ($1, $2)`, user.Id, user.Username)

	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	if err = tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (s *Storage) GetUser(ctx context.Context, id uuid.UUID) (*models.User, error) {
	var user models.User
	err := s.db.QueryRow(ctx, `SELECT 
    id, 
    firstname, 
    lastname,
    avatar_url,
	website,
	location,
	birth_date,
	gender,
	telephone
FROM user_profiles WHERE id=$1`, id).Scan(
		&user.Id,
		&user.Username,
		&user.Firstname,
		&user.Lastname,
		&user.AvatarUrl,
		&user.Website,
		&user.Location,
		&user.BirthDate,
		&user.Gender,
		&user.Telephone,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domainErrors.ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to load user from DB: %w", err)
	}

	return &user, nil
}

func (s *Storage) UpdateUser(ctx context.Context, user *models.User) error {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	defer tx.Rollback(ctx)

	var exists bool
	err = tx.QueryRow(ctx, `SELECT EXISTS (SELECT 1 FROM user_profiles WHERE id=$1)`, user.Id).Scan(&exists)

	if err != nil {
		return fmt.Errorf("failed to check if user exists: %w", err)
	}

	if !exists {
		return fmt.Errorf("user does not exist")
	}

	query := `UPDATE user_profiles SET`
	args := []interface{}{}
	paramCount := 1

	updateFields := map[string]interface{}{
		"firstname":  user.Firstname,
		"lastname":   user.Lastname,
		"avatar_url": user.AvatarUrl,
		"website":    user.Website,
		"location":   user.Location,
		"bio":        user.Bio,
		"birthdate":  user.BirthDate,
		"gender":     user.Gender,
		"telephone":  user.Telephone,
	}

	for fieldName, fieldValue := range updateFields {
		if fieldValue != nil {
			query += fmt.Sprintf(" %s = $%d", fieldName, paramCount)
			args = append(args, fieldValue)
			paramCount++
		}
	}

	query += " WHERE id = $" + strconv.Itoa(paramCount)
	args = append(args, user.Id)

	_, err = tx.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
