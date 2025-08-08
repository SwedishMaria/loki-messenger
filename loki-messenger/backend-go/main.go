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

	ws := api.Group("/ws")
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

	go handlers.HandleMessages()

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}