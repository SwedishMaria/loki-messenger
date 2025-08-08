package models

import (
	"context"
	"time"
	"github.com/loki-messenger/backend-go/database"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int       `json:"id"`
	Username     string    `json:"username"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func CreateUser(username, password string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &User{
		Username:     username,
		PasswordHash: string(hashedPassword),
	}

	query := `
		INSERT INTO users (username, password_hash)
		VALUES ($1, $2)
		RETURNING id, created_at, updated_at
	`
	err = database.DB.QueryRow(context.Background(), query, user.Username, user.PasswordHash).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserByUsername(username string) (*User, error) {
	user := &User{}
	query := `SELECT id, username, password_hash, created_at, updated_at FROM users WHERE username = $1`
	err := database.DB.QueryRow(context.Background(), query, username).Scan(&user.ID, &user.Username, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	return err == nil
}