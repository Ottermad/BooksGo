package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

// Function to connect to database
func NewDB() *sql.DB {
	db, err := sql.Open("sqlite3", "sqlite.db")

	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS TBL_BOOK (ID INTEGER PRIMARY KEY AUTOINCREMENT, TITLE VARCHAR(100), AUTHOR VARCHAR(100))")

	if err != nil {
		panic(err)
	}

	return db

}

// Function to insert new book
func InsertBook(db *sql.DB, title string, author string) (bool, string) {
	_, err := db.Exec("INSERT INTO TBL_BOOK(TITLE, AUTHOR) VALUES (?, ?)", title, author)

	if err != nil {
		panic(err)
	}

	return true, ""
}

// Function to read book
func ReadBooks(db *sql.DB) map[string]map[string]string {
	rows, err := db.Query("SELECT * FROM TBL_BOOK")

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	data := make(map[string]map[string]string)

	for rows.Next() {
		var id, title, author string
		if err := rows.Scan(&id, &title, &author); err != nil {
			panic(err)
		}

		data[id] = map[string]string{
			"author": author,
			"title":  title,
		}
		fmt.Println("The title of this book is ", title)
	}

	if err = rows.Err(); err != nil {
		panic(err)
	}

	return data
}
