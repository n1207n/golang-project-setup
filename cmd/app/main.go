package main

import (
	_ "github.com/lib/pq"
	app "github.com/n1207n/golang-project-scaffold/internal/app"
	"log"
	"os"
)

func main() {
	// Read environment variables
	dbConfig := app.DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
	}

	redisConfig := app.RedisConfig{
		Host: os.Getenv("REDIS_HOST"),
		Port: os.Getenv("REDIS_PORT"),
	}

	// Wait for the database to be ready
	if err := app.WaitForDB(dbConfig); err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	server, err := app.NewServer(dbConfig, redisConfig)
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}
	defer server.Close()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s", port)
	if err := server.Start(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
