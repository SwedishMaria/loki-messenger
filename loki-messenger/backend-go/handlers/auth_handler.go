package handlers

import (
	"context"
	"loki-messenger/database"
	"loki-messenger/models"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Register(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to hash password")
	}

	_, err = database.Conn.Exec(context.Background(), "INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, string(hashedPassword))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create user")
	}

	return c.JSON(http.StatusCreated, "User created successfully")
}

func Login(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	storedUser := new(models.User)
	err := database.Conn.QueryRow(context.Background(), "SELECT id, username, password FROM users WHERE username=$1", user.Username).Scan(&storedUser.ID, &storedUser.Username, &storedUser.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid credentials")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = storedUser.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to generate token")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}