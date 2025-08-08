package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/loki-messenger/backend-go/database"
	"github.com/loki-messenger/backend-go/handlers"
	"golang.org/x/time/rate"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}
	if err := database.Connect(); err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer database.Close()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api := e.Group("/api")
	auth := api.Group("/auth")

	auth.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(5))))

	auth.POST("/register", handlers.Register)
	auth.POST("/login", handlers.Login)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}