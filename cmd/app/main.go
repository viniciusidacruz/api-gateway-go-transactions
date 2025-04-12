package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/viniciusidacruz/api-gateway-go-transactions/internal/modules/account/repository"
	"github.com/viniciusidacruz/api-gateway-go-transactions/internal/modules/account/services"
	"github.com/viniciusidacruz/api-gateway-go-transactions/internal/web/server"
)

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		getEnv("DB_HOST", "db"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_NAME", "gateway"),
		getEnv("DB_SSLMODE", "disable"),
	)

	db, err := sql.Open("postgres", connStr)
	
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	defer db.Close()

	accountRepository := repository.NewAccountRepository(db)
	accountService := services.NewAccountService(accountRepository)

	port := getEnv("HTTP_PORT", "8080")
	server := server.NewServer(port, accountService)

	server.ConfigureRoutes()
	server.Start()

	if err := server.Start(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

