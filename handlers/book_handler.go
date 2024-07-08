package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"

	"crud-library-book/model"
	db "crud-library-book/utils"

	"github.com/gorilla/mux"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	json.NewDecoder(r.Body).Decode(&book)

	query := `INSERT INTO books (uuid, title, isbn, author, publisher, year, category, location, eksemplar)
            VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING uuid`
	err := db.DB.QueryRowContext(context.Background(), query, book.UUID, book.Title, book.ISBN, book.Author, book.Publisher, book.Year, book.Category, book.Location, book.Eksemplar).Scan(&book.UUID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.QueryContext(context.Background(), "SELECT uuid, title, isbn, author, publisher, year, category, location, eksemplar FROM books")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var books []model.Book
	for rows.Next() {
		var book model.Book
		err := rows.Scan(&book.UUID, &book.Title, &book.ISBN, &book.Author, &book.Publisher, &book.Year, &book.Category, &book.Location, &book.Eksemplar)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		books = append(books, book)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func GetOneBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]

	var book model.Book
	query := `SELECT uuid, title, isbn, author, publisher, year, category, location, eksemplar 
            FROM books WHERE uuid = $1`
	err := db.DB.QueryRowContext(context.Background(), query, uuid).Scan(&book.UUID, &book.Title, &book.ISBN, &book.Author, &book.Publisher, &book.Year, &book.Category, &book.Location, &book.Eksemplar)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Book not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]

	var book model.Book
	json.NewDecoder(r.Body).Decode(&book)
	book.UUID = uuid

	query := `UPDATE books SET title = $1, isbn = $2, author = $3, publisher = $4, year = $5, category = $6, location = $7, eksemplar = $8 WHERE uuid = $9`
	_, err := db.DB.ExecContext(context.Background(), query, book.Title, book.ISBN, book.Author, book.Publisher, book.Year, book.Category, book.Location, book.Eksemplar, book.UUID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]

	query := "DELETE FROM books WHERE uuid = $1"
	_, err := db.DB.ExecContext(context.Background(), query, uuid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
