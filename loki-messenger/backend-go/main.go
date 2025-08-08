package main

import (
	"loki-messenger/database"
	"loki-messenger/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Initialize Database
	database.InitDB()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))

	// Routes
	api := e.Group("/api")
	auth := api.Group("/auth")
	auth.POST("/register", handlers.Register)
	auth.POST("/login", handlers.Login)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}