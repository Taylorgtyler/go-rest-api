package main

import (
	"log"

	"github.com/Taylorgtyler/go-rest-api/cmd/pkg/database"
	"github.com/Taylorgtyler/go-rest-api/cmd/pkg/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	r := gin.Default() // Creates a router without any middleware by default

	// Set up your routes
	handlers.SetupRoutes(r, db)

	// Start the server
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
