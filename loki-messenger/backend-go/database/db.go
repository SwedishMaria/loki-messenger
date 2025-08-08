package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	DB *pgxpool.Pool
)

func Connect() error {
	config, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		return fmt.Errorf("failed to parse database config: %w", err)
	}

	db, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	DB = db
	return nil
}

func Close() {
	if DB != nil {
		DB.Close()
	}
}