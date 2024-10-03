package main

import (
	"bookstore/db"     // Adjust according to your folder structure
	"bookstore/routes" // Ensure the import path is correct
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the database
	db.ConnectDatabase() // Ensure this function doesn't return a value
	db.SeedBooks()      // Seed the database with initial data

	// Set up Gin router
	r := gin.Default()

	// Setup the routes for the application
	routes.SetupRoutes(r)

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Unable to start the server:", err)
	}
}
