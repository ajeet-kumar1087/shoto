package database

import (
	"log"
	"shoto/app/user"
)

func RunMigrations() {
	log.Println("Running migrations...")
	err := DB.AutoMigrate(&user.User{})
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
}
