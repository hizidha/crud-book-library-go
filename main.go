package main

import (
	"log"
	"net/http"

	handler "crud-library-book/handlers" // Sesuaikan dengan path package handler Anda
	db "crud-library-book/utils"         // Sesuaikan dengan path package db Anda

	"github.com/gorilla/mux"
)

func main() {
	// Inisialisasi koneksi ke database
	err := db.InitDB()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer db.DB.Close()

	// Membuat router menggunakan Gorilla Mux
	r := mux.NewRouter()

	// Menambahkan handler untuk rute-rute HTTP
	r.HandleFunc("/books", handler.CreateBook).Methods("POST")
	r.HandleFunc("/books", handler.GetBooks).Methods("GET")
	r.HandleFunc("/books/{uuid}", handler.GetOneBook).Methods("GET")
	r.HandleFunc("/books/{uuid}", handler.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{uuid}", handler.DeleteBook).Methods("DELETE")

	// Menjalankan server HTTP
	log.Fatal(http.ListenAndServe(":8080", r))
}
