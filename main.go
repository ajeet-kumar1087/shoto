package main

import (
	"log"
	"shoto/pkg/database"
)

func main() {
	database.ConnectDB()
	log.Println("DB connected successfully")
}
