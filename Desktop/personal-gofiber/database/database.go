package database

import (
    "database/sql"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
    var err error
    DB, err = sql.Open("sqlite3", "./blog.db")
    if err != nil {
        log.Fatal(err)
    }

    query := `
    CREATE TABLE IF NOT EXISTS posts (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        content TEXT NOT NULL
    );`
    _, err = DB.Exec(query)
    if err != nil {
        log.Fatal(err)
    }
}
