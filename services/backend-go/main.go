package main

import (
	"loki-messenger/database"
	"loki-messenger/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Initialize the database connection.
	database.InitDB()

	// --- Middleware --- //

	// Logger logs information about each HTTP request.
	e.Use(middleware.Logger())

	// Recover recovers from panics anywhere in the chain.
	e.Use(middleware.Recover())

	// RateLimiter limits the number of requests from a single IP.
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))

	// --- Routes --- //

	api := e.Group("/api")

	// Authentication routes
	auth := api.Group("/auth")
	auth.POST("/register", handlers.Register)
	auth.POST("/login", handlers.Login)

	// WebSocket routes
	ws := api.Group("/ws")
	// This middleware extracts the JWT from the query param and validates it.
	ws.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.QueryParam("token")
			if token == "" {
				return echo.ErrUnauthorized
			}
			return next(c)
		}
	})
	ws.GET("", handlers.HandleConnections)

	// Start a goroutine to handle incoming WebSocket messages.
	go handlers.HandleMessages()

	// Start the server on port 8080.
	e.Logger.Fatal(e.Start(":8080"))
}