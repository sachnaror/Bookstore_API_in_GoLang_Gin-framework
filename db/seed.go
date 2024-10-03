package db

import "log"

func SeedBooks() {
    seedSQL := `INSERT INTO books (title, author, publish_date, description) VALUES
        ('Book Title 1', 'Author 1', '2024-01-01', 'Description 1'),
        ('Book Title 2', 'Author 2', '2024-01-02', 'Description 2'),
        ('Book Title 3', 'Author 3', '2024-01-03', 'Description 3');`

    if _, err := DB.Exec(seedSQL); err != nil {
        log.Printf("Could not seed books: %v", err)
    }
}
