package app

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type QueryService struct {
	db *sql.DB
}

func NewQueryService(config DBConfig) (*QueryService, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Name)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return &QueryService{db: db}, nil
}

func (q *QueryService) Close() error {
	return q.db.Close()
}

// Add more methods for database operations here
