package main

import (
	"bookstore/db"
	"bookstore/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
    db.ConnectDatabase()
    db.SeedBooks()

    r := gin.Default()

    // Load HTML templates
    r.LoadHTMLGlob("templates/*") // Adjust this path as needed

    routes.SetupRoutes(r)

    if err := r.Run(":8080"); err != nil {
        log.Fatal("Unable to start the server:", err)
    }
}
