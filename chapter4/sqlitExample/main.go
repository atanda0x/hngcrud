package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// Book is a placeholder for book
type Book struct {
	id     int
	name   string
	author string
}

func dbOperations(db *sql.DB) {
	// Create
	schema, _ := db.Prepare("INSERT INTO books (name, author, isbn) VALUES (?, ?, ?)")
	schema.Exec("A Tale of Two Cities", "Charles Dickens", 140430547)
	log.Println("Inserted the book info into db!!!")

	// Read
	rows, _ := db.Query("SELECT id, name, author FROM books")
	var tempBook Book
	for rows.Next() {
		rows.Scan(&tempBook.id, &tempBook.name, &tempBook.author)
		log.Printf("ID: %d, Book:%s, Author:%s", tempBook.id, tempBook.name, tempBook.author)
	}

	// Update
	schem, _ := db.Prepare("UPDATE books SET name=? WHERE id=?")
	schem.Exec("The Tale of Two Cities", 1)
	log.Println("Successful updated the book in db!!!!!")

	// Delete
	del, _ := db.Prepare("DELET FROM books WHERE id=?")
	del.Exec(1)
	log.Println("Successfully deleted the book in db!!!!")
}

func main() {
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		log.Println(err)
	}

	// create table
	schema, err := db.Prepare("CREATE TABLE IF NOT EXISTS books(id INTEGER PRIMARY KEY, isbn INTEGER, author VARCHAR(64), name VARCHAR(64) NULL)")
	if err != nil {
		log.Println("Errr in Creating table")
	} else {
		log.Println("Successfully created table books!!")
	}
	schema.Exec()
	dbOperations(db)
}
