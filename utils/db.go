package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB // Variabel global untuk menyimpan koneksi database

func InitDB() error {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file: %v", err)
	}

	// Read database connection details from environment variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// String koneksi
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Membuka koneksi ke database
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}

	// Menguji koneksi
	err = db.Ping()
	if err != nil {
		db.Close()
		return fmt.Errorf("error connecting to database: %v", err)
	}

	log.Println("Connected to PostgreSQL!")

	// Menetapkan koneksi ke variabel global
	DB = db

	return nil
}
