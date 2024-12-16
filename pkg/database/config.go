package utils

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sql.DB

// ConnectDB menginisialisasi koneksi database dan menyimpannya di variabel global DB
func ConnectDB(user, password, dbname string, host string, port string) (*sql.DB, error) {
	if DB != nil {
		return DB, nil
	}

	psqlInfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", user, password, dbname, host, port)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	// Test koneksi
	if err := db.Ping(); err != nil {
		return nil, err
	}

	DB = db
	return DB, nil
}

// CloseDB menutup koneksi database
func CloseDB() {
	if DB != nil {
		err := DB.Close()
		if err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}
}
