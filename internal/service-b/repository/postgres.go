package repository

import (
	"database/sql"
	"fmt"
	"os"
	"transaction-worker/internal/common/config"

	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func (s *Storage) DB() *sql.DB {
	return s.db
}

func Connect(cfg config.Datebase) (*Storage, error) {
	const op = "storage.postgre.New"

	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		getEnv("DB_HOST", cfg.Host),
		getEnv("DB_PORT", cfg.Port),
		getEnv("DB_USER", cfg.User),
		getEnv("DB_PASSWORD", cfg.Password),
		getEnv("DB_NAME", cfg.Dbname),
	)

	db, err := sql.Open("postgres", sqlInfo)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
