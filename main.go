package main

import (
	"log"

	"github.com/labstack/echo"
)

func main() {
	log.Println("Starting the application....")
	e := echo.New()
	e.Logger.Fatal(e.Start(":9090"))

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello go")
	})

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "ok", "message": "Health is okay okay ok"})
	})

	// create DB connection
	// db := database.ConnectDB()
	// sqlDB, err := db.DB()
	// if err != nil {
	// 	log.Fatalf("Failed to get database connection: %v", err)
	// }

	// defer sqlDB.Close()

	// log.Println("DB connected successfully")
	// Initialize the repository
	// userRepo := user.NewUserRepository(db)

	// // Initialize the service
	// userService := user.NewUserService(userRepo)

	// // Initialize the handler
	// userHandler := user.NewUserHandler(userService)

	// // // Set up the router
	// user.Routes(e, userHandler)

	// // Start the server
	// e.Logger.Fatal(e.Start(":8081"))
}
