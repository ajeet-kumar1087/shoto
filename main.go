package main

import (
	"log"
	"shoto/app/user"
	"shoto/pkg/database"

	"github.com/labstack/echo"
)

func main() {
	log.Println("Starting the application....")
	e := echo.New()

	// create DB connection
	db := database.ConnectDB()
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database connection: %v", err)
	}

	defer sqlDB.Close()

	log.Println("DB connected successfully")
	// Initialize the repository
	userRepo := user.NewRepository(db)

	// // Initialize the service
	userService := user.NewService(userRepo)

	// // Initialize the handler
	userHandler := user.NewHandler(userService)

	// // // Set up the router
	user.Routes(e, userHandler)

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
