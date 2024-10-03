package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnectDatabase() {
    var err error
    DB, err = sql.Open("sqlite3", "./books.db") // Make sure the path is correct
    if err != nil {
        log.Fatal(err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatal(err)
    }

    // Create table if it doesn't exist
    createTableSQL := `CREATE TABLE IF NOT EXISTS books (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        author TEXT NOT NULL,
        publish_date TEXT NOT NULL,
        description TEXT
    );`

    if _, err = DB.Exec(createTableSQL); err != nil {
        log.Fatal(err)
    }
}
