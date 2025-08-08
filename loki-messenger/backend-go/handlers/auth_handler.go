package handlers

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/loki-messenger/backend-go/models"
)

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(c echo.Context) error {
	req := new(AuthRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	if req.Username == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Username and password are required"})
	}

	user, err := models.CreateUser(req.Username, req.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create user"})
	}

	return c.JSON(http.StatusCreated, user)
}

func Login(c echo.Context) error {
	req := new(AuthRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	user, err := models.GetUserByUsername(req.Username)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid credentials"})
	}

	if !user.CheckPassword(req.Password) {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid credentials"})
	}

	claims := &jwt.RegisteredClaims{
		Subject:   user.Username,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create token"})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}