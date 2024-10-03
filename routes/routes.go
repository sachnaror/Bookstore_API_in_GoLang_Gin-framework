package routes

import (
	"bookstore/controllers" // Ensure this import path is correct

	"github.com/gin-gonic/gin"
)

// SetupRoutes sets up the routing for the application
func SetupRoutes(r *gin.Engine) {
	r.GET("/books", controllers.GetBooks)      // To get all books
	r.GET("/books/:id", controllers.GetBook)   // To get a specific book by ID
	r.POST("/books", controllers.CreateBook)    // To create a new book
	r.PUT("/books/:id", controllers.EditBook)   // To edit an existing book
	r.DELETE("/books/:id", controllers.DeleteBook) // To delete a book
	r.POST("/books/:id/edit", controllers.EditBook) // Optional: to handle edit via POST
}
