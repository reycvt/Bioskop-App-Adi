package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {

	host := "localhost"
	port := 5432
	user := "testgin"
	password := "testgin123"
	dbname := "bioskop_adi"

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Gagal koneksi ke database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Database tidak merespon: %v", err)
	}

	DB = db
	fmt.Println("Berhasil konek ke PostgreSQL")
}
