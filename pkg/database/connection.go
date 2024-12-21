package database

import (
	"fmt"
	"log"
	"net/url"
	"shoto/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	config.LoadEnv()

	// Use net/url to construct the DSN
	dbURL := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(config.GetEnv("DB_USER", "postgres"), config.GetEnv("DB_PASSWORD", "")),
		Host:   fmt.Sprintf("%s:%s", config.GetEnv("DB_HOST", "localhost"), config.GetEnv("DB_PORT", "5432")),
		Path:   config.GetEnv("DB_NAME", "marketplace"),
	}

	// Add query parameters
	query := dbURL.Query()
	query.Add("sslmode", config.GetEnv("DB_SSL_MODE", "disable"))
	dbURL.RawQuery = query.Encode()

	// Connect to the database
	var err error
	DB, err = gorm.Open(postgres.Open(dbURL.String()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	return DB
}
