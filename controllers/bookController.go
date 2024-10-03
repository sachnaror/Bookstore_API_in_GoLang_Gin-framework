package controllers

import (
	"bookstore/db"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Define the struct for the book
type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	PublishDate string `json:"publish_date"`
	Description string `json:"description"`
}

// GetBooks handles the GET request to retrieve all books.
func GetBooks(c *gin.Context) {
	rows, err := db.DB.Query("SELECT id, title, author, publish_date, description FROM books")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch books"})
		return
	}
	defer rows.Close()

	var books []Book // Use the defined Book struct here

	for rows.Next() {
		var book Book // Use the same struct type for each book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.PublishDate, &book.Description); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not scan book"})
			return
		}
		books = append(books, book) // Append the book to the slice
	}

	// Beautify JSON response
	prettyJSON, err := json.MarshalIndent(books, "", "  ")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate JSON"})
		return
	}

	// Set the Content-Type header and return the pretty JSON
	c.Data(http.StatusOK, "application/json", prettyJSON) // Send the pretty JSON response
}
// GetBook handles the GET request to retrieve a specific book by ID.
func GetBook(c *gin.Context) {
	id := c.Param("id")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	var book Book
	err = db.DB.QueryRow("SELECT id, title, author, publish_date, description FROM books WHERE id = ?", bookID).
		Scan(&book.ID, &book.Title, &book.Author, &book.PublishDate, &book.Description)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		log.Printf("Error fetching book: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch book"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// CreateBook adds a new book to the database.
func CreateBook(c *gin.Context) {
	var newBook Book

	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	_, err := db.DB.Exec(`INSERT INTO books (title, author, publish_date, description) VALUES (?, ?, ?, ?)`,
		newBook.Title, newBook.Author, newBook.PublishDate, newBook.Description)
	if err != nil {
		log.Printf("Error adding book: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not add book"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Book created", "book": newBook})
}

// EditBook updates an existing book in the database.
func EditBook(c *gin.Context) {
	id := c.Param("id")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	var updatedBook Book
	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	_, err = db.DB.Exec(`UPDATE books SET title = ?, author = ?, publish_date = ?, description = ? WHERE id = ?`,
		updatedBook.Title, updatedBook.Author, updatedBook.PublishDate, updatedBook.Description, bookID)
	if err != nil {
		log.Printf("Error updating book: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book updated", "book": updatedBook})
}

// DeleteBook removes a book from the database.
func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	_, err = db.DB.Exec(`DELETE FROM books WHERE id = ?`, bookID)
	if err != nil {
		log.Printf("Error deleting book: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
