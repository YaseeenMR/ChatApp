package repository

import (
	"database/sql"

	"github.com/yourusername/chat-app/backend/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`
	err = r.db.QueryRow(query, user.Name, user.Email, string(hashedPassword)).Scan(&user.ID)
	return err
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, name, email, password FROM users WHERE email = $1`
	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(userID uint, updateData *models.UpdateProfileRequest) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	if updateData.Name != "" {
		_, err = tx.Exec("UPDATE users SET name = $1 WHERE id = $2", updateData.Name, userID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if updateData.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updateData.Password), bcrypt.DefaultCost)
		if err != nil {
			tx.Rollback()
			return err
		}
		_, err = tx.Exec("UPDATE users SET password = $1 WHERE id = $2", hashedPassword, userID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (r *UserRepository) GetUserByID(userID uint) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, name, email FROM users WHERE id = $1`
	err := r.db.QueryRow(query, userID).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
