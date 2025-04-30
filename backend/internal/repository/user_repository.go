package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/yourusername/chat-app/backend/internal/models"
	"golang.org/x/crypto/bcrypt"
)

// Define custom error types for better error handling
var (
	ErrDuplicateEmail    = errors.New("email already exists")
	ErrUserNotFound      = errors.New("user not found")
	ErrDatabaseOperation = errors.New("database operation failed")
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	user.Email = strings.ToLower(user.Email)
	if r.db == nil {
		return fmt.Errorf("%w: database connection is nil", ErrDatabaseOperation)
	}

	// Check email existence first
	var exists bool
	err := r.db.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)",
		user.Email,
	).Scan(&exists)

	if err != nil {
		return fmt.Errorf("%w: %v", ErrDatabaseOperation, err)
	}
	if exists {
		return ErrDuplicateEmail
	}

	// Insert new user with context timeout
	query := `INSERT INTO users (name, email, password, created_at, updated_at) 
              VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err = r.db.QueryRow(
		query,
		user.Name,
		user.Email,
		user.Password, // Password should already be hashed
		time.Now(),
		time.Now(),
	).Scan(&user.ID)

	if err != nil {
		return fmt.Errorf("%w: %v", ErrDatabaseOperation, err)
	}
	return nil
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	email = strings.ToLower(email)
	user := &models.User{}
	query := `SELECT id, name, email, password, created_at, updated_at 
              FROM users WHERE email = $1`
	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("%w: %v", ErrDatabaseOperation, err)
	}
	return user, nil
}

func (r *UserRepository) GetUserByID(id uint) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, name, email, password, created_at, updated_at 
              FROM users WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("%w: %v", ErrDatabaseOperation, err)
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(userID uint, updateData *models.UpdateProfileRequest) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("%w: %v", ErrDatabaseOperation, err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	if updateData.Name != "" {
		_, err = tx.Exec(
			"UPDATE users SET name = $1, updated_at = $2 WHERE id = $3",
			updateData.Name,
			time.Now(),
			userID,
		)
		if err != nil {
			return fmt.Errorf("%w: %v", ErrDatabaseOperation, err)
		}
	}

	if updateData.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword(
			[]byte(updateData.Password),
			bcrypt.DefaultCost,
		)
		if err != nil {
			return fmt.Errorf("failed to hash password: %w", err)
		}

		_, err = tx.Exec(
			"UPDATE users SET password = $1, updated_at = $2 WHERE id = $3",
			hashedPassword,
			time.Now(),
			userID,
		)
		if err != nil {
			return fmt.Errorf("%w: %v", ErrDatabaseOperation, err)
		}
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("%w: %v", ErrDatabaseOperation, err)
	}
	return nil
}
