package models

// Book represents a book structure
type Book struct {
    ID          int    `json:"id"`
    Title       string `json:"title"`
    Author      string `json:"author"`
    PublishDate string `json:"publish_date"`
    Description string `json:"description"`
}
